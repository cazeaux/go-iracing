package iracing

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/cazeaux/go-iracing/cmd/config"
	"github.com/cazeaux/go-iracing/cmd/store"
	"github.com/cazeaux/go-iracing/pkg/iracing"
	"github.com/cazeaux/go-iracing/pkg/types"
)

type IRacingService struct {
	logger      *slog.Logger
	client      *iracing.Client
	dataStore   *store.DataStore
	assetsCache *IracingAssetsCache
}

func (ir *IRacingService) GetUserLastResults(ctx context.Context, u *config.User) ([]RaceResult, error) {
	tech := &techData{
		service: ir,
		ctx:     ctx,
	}
	state := &stateData{}
	out := &outData{}
	in := &inData{
		user: u,
	}

	f := getUserCachedData

	for {
		f = f(tech, state, in, out)
		if f == nil {
			break
		}
	}

	return nil, nil
}

type techData struct {
	service *IRacingService
	ctx     context.Context
}

type stateData struct {
	recentRaces   []types.RecentRace
	subsessions   map[int]*types.ResultsGetResp
	userCacheData *store.Data
}

type outData struct {
	results []*RaceResult
	err     error
}

type inData struct {
	user *config.User
}

type ResultStateFunc func(tech *techData, state *stateData, in *inData, out *outData) ResultStateFunc

func getUserCachedData(tech *techData, state *stateData, in *inData, out *outData) ResultStateFunc {
	userCacheData, err := store.Get(tech.service.dataStore, in.user)
	if err != nil {
		tech.service.logger.Error("failed to retrieve user cache data", "user", in.user.ID, "err", err)
		out.err = err

		return nil
	}
	state.userCacheData = userCacheData
	return getMemberRecentRaces
}

func getMemberRecentRaces(tech *techData, state *stateData, in *inData, out *outData) ResultStateFunc {
	req := &types.StatsMemberRecentRacesReq{
		CustID: in.user.ID,
	}

	memberResults, _, err := tech.service.client.Stats.MemberRecentResult(tech.ctx, req)
	if err != nil {
		tech.service.logger.Error("failed to retrieve member recent results", "user", in.user.ID, "err", err)
		out.err = err

		return nil
	}

	// Take only the
	if state.userCacheData.LastSubSessionID == 0 {
		state.recentRaces = memberResults.Races[len(memberResults.Races)-1:]
	} else {
		for _, r := range memberResults.Races {
			if r.SubsessionID <= state.userCacheData.LastSubSessionID {
				continue
			}
			state.recentRaces = append(state.recentRaces, r)
		}
	}

	return getSubsessions
}

func getSubsessions(tech *techData, state *stateData, in *inData, out *outData) ResultStateFunc {
	for _, r := range state.recentRaces {
		req := types.ResultsGetReq{
			SubsessionID: r.SubsessionID,
		}
		subsession, _, err := tech.service.client.Results.Get(tech.ctx, &req)
		if err != nil {
			tech.service.logger.Error("failed to retrieve subsession", "subsession", r.SubsessionID, "err", err)
			out.err = err

			return nil
		}

		state.subsessions[r.SubsessionID] = subsession
	}

	return prepareRaceResults
}

func prepareRaceResults(tech *techData, state *stateData, in *inData, out *outData) ResultStateFunc {
	for _, race := range state.recentRaces {
		subsession, ok := state.subsessions[race.SubsessionID]
		if !ok {
			tech.service.logger.Error("could not find subsession", "subsession", race.SubsessionID)
			continue
		}
		res, err := buildRaceResult(&race, subsession, in.user, tech.service.assetsCache)
		if err != nil {
			tech.service.logger.Error("faild to build results", "subsession", race.SubsessionID, "user", in.user.ID)
			continue
		}
		out.results = append(out.results, res)
	}
	return updateUserCachedData

}

func updateUserCachedData(tech *techData, state *stateData, in *inData, out *outData) ResultStateFunc {
	data := &store.Data{
		LastSubSessionID:    state.recentRaces[len(state.recentRaces)-1].SubsessionID,
		LastResultTimestamp: state.recentRaces[len(state.recentRaces)-1].SessionStartTime,
	}

	tech.service.dataStore.WriterChannel <- store.Message{
		User: in.user,
		Data: data,
	}

	return nil
}

func buildRaceResult(race *types.RecentRace, subsession *types.ResultsGetResp, user *config.User, assetsCache *IracingAssetsCache) (*RaceResult, error) {
	raceResult := RaceResult{
		FieldSize: getSessionFieldSize(subsession, types.SessionTypeQualification),
	}

	raceSessionResults := getResultForSessionType(subsession, user, types.SessionTypeRace)
	if raceSessionResults == nil {
		return nil, fmt.Errorf("race session not found")
	}

	qualifSessionResults := getResultForSessionType(subsession, user, types.SessionTypeQualification)
	if qualifSessionResults == nil {
		return nil, fmt.Errorf("race session not found")
	}

	// irating, srating
	raceResult.NewIRating = raceSessionResults.NewiRating
	raceResult.OldIRating = raceSessionResults.OldiRating
	raceResult.NewSRating = raceSessionResults.NewSubLevel
	raceResult.OldIRating = raceSessionResults.OldSubLevel

	// Race results
	raceResult.RacePos = raceSessionResults.FinishPosition
	raceResult.BestLapTime = raceSessionResults.BestLapTime
	raceResult.ChamPoints = raceSessionResults.ChampPoints

	// Quali results
	raceResult.QualyPos = qualifSessionResults.FinishPosition
	raceResult.QualiLapTime = qualifSessionResults.BestQualLapTime

	// Splits
	_, splitNumber := getSessionSplit(subsession)
	raceResult.Splits = len(subsession.SessionSplits)
	raceResult.Split = splitNumber

	// Pilot, track, car, series
	raceResult.CarID = raceSessionResults.CarID
	raceResult.CarName = raceSessionResults.CarName
	raceResult.TrackID = race.Track.TrackID
	raceResult.TrackName = race.Track.TrackName
	raceResult.PilotDisplayName = raceSessionResults.DisplayName
	raceResult.SeriesID = race.SeriesID
	raceResult.SeriesName = race.SeriesName

	// Incidents
	raceResult.Incidents = raceSessionResults.Incidents

	// SOF, field size
	raceResult.SOF = race.StrengthOfField
	raceResult.FieldSize = getSessionFieldSize(subsession, types.SessionTypeQualification)

	// Start time
	raceResult.StartTime = race.SessionStartTime

	// Multi class
	raceResult.IsMultiClass = len(subsession.CarClasses) > 1
	if raceResult.IsMultiClass {
		sof, fieldSize := getSessionFieldForClass(subsession, types.SessionTypeRace, raceSessionResults.CarClassID)

		raceResult.MultiClass = &MultiClass{
			FieldSize: fieldSize,
			SOF:       sof,
			RacePos:   raceSessionResults.FinishPositionInClass,
			QualyPos:  qualifSessionResults.FinishPositionInClass,
		}
	}

	// Weather
	if w := GetSessionWeatherResults(subsession, types.SessionTypeQualification); w != nil {
		raceResult.QualyRain = w.PrecipMm > 0
	}
	if w := GetSessionWeatherResults(subsession, types.SessionTypeRace); w != nil {
		raceResult.RaceRain = w.PrecipMm > 0
	}

	// Logos
	raceResult.CarLogo = (*assetsCache.Cars)[string(raceResult.CarID)].Logo
	raceResult.SeriesLogo = (*assetsCache.Cars)[string(raceResult.SeriesID)].Logo

	return &raceResult, nil
}

func getResultForSessionType(subsession *types.ResultsGetResp, user *config.User, sessionType types.SessionType) *types.SubsessionSessionResultsPilots {
	for _, s := range subsession.SessionResults {
		if s.SimsessionType != int(sessionType) {
			continue
		}
		for _, r := range s.Results {
			if r.CustID == user.ID {
				return &r
			}
		}
	}
	return nil
}

func getSessionFieldSize(subsession *types.ResultsGetResp, sessionType types.SessionType) int {
	for _, s := range subsession.SessionResults {
		if s.SimsessionType == int(sessionType) {
			return len(s.Results)
		}
	}
	return 0
}

func getSessionFieldForClass(subsession *types.ResultsGetResp, sessionType types.SessionType, carClassID int) (sof, fieldSize int) {
	fieldSize = 0
	iratings := 0
	for _, s := range subsession.SessionResults {
		if s.SimsessionType != int(sessionType) {
			continue
		}
		for _, r := range s.Results {
			if r.CarClassID == carClassID {
				fieldSize++
				iratings += r.NewiRating
			}
		}
	}

	return iratings / fieldSize, fieldSize
}

func getSessionSplit(subsession *types.ResultsGetResp) (*types.SubsessionSplit, int) {
	for n, s := range subsession.SessionSplits {
		if s.SubsessionID == subsession.SubsessionID {
			return &s, n
		}
	}
	return nil, 0
}

func GetSessionWeatherResults(subsession *types.ResultsGetResp, sessionType types.SessionType) *types.SessionWeatherResult {
	for _, s := range subsession.SessionResults {
		if s.SimsessionType == int(sessionType) {
			return &s.WeatherResult
		}
	}
	return nil
}

// TODO: get split
