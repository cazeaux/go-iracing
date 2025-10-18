package iracing

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cazeaux/go-iracing/pkg/types"
	"github.com/google/go-querystring/query"
)

type StatsService struct{ client *Client }

func (s *StatsService) MemberRecentResult(ctx context.Context, req *types.StatsMemberRecentRacesReq) (*types.StatsMemberRecentRacesResp, *http.Response, error) {
	path := "/data/stats/member_recent_races"
	var infoResp types.StatsMemberRecentRacesResp

	params := url.Values{}
	params.Add("cust_id", strconv.Itoa(req.CustID))

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *StatsService) MemberCareer(ctx context.Context, req *types.StatsMemberCareerReq) (*types.StatsMemberCareerResp, *http.Response, error) {
	path := "/data/stats/member_career"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.StatsMemberCareerResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *StatsService) MemberDivision(ctx context.Context, req *types.StatsMemberDivisionReq) (*types.StatsMemberDivisionResp, *http.Response, error) {
	path := "/data/stats/member_division"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.StatsMemberDivisionResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *StatsService) MemberSummary(ctx context.Context, req *types.StatsMemberSummaryReq) (*types.StatsMemberSummaryResp, *http.Response, error) {
	path := "/data/stats/member_summary"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.StatsMemberSummaryResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *StatsService) MemberYearly(ctx context.Context, req *types.StatsMemberYearlyReq) (*types.StatsMemberYearlyResp, *http.Response, error) {
	path := "/data/stats/member_yearly"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.StatsMemberYearlyResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *StatsService) MemberBests(ctx context.Context, req *types.StatsMemberBestsReq) (*types.StatsMemberBestsResp, *http.Response, error) {
	path := "/data/stats/member_bests"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.StatsMemberBestsResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *StatsService) MemberRecap(ctx context.Context, req *types.StatsMemberRecapReq) (*types.StatsMemberRecapResp, *http.Response, error) {
	path := "/data/stats/member_recap"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.StatsMemberRecapResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *StatsService) WorldRecords(ctx context.Context, req *types.StatsWorldRecordsReq) ([]types.StatsWorldRecordsResp, *http.Response, error) {
	path := "/data/stats/world_records"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp []types.StatsWorldRecordsResp
	respData, err := GetRessourceChunks(ctx, s.client, path, params, req.ChunkReq.Rows, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return infoResp, respData, nil
}

func (s *StatsService) SeasonQualifyResults(ctx context.Context, req *types.StatsSeasonQualifyResultsReq) ([][]string, *http.Response, error) {
	path := "/data/stats/season_qualify_results"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp [][]string
	respData, err := s.client.getRessourceDataCsvURL(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return infoResp, respData, nil
}

func (s *StatsService) SeasonDriverStandings(ctx context.Context, req *types.StatsSeasonDriverStandingsReq) ([][]string, *http.Response, error) {
	path := "/data/stats/season_driver_standings"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp [][]string
	respData, err := s.client.getRessourceDataCsvURL(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return infoResp, respData, nil
}

func (s *StatsService) SeasonSupersessionStandings(ctx context.Context, req *types.StatsSeasonSupersessionStandingsReq) ([][]string, *http.Response, error) {
	path := "/data/stats/season_supersession_standings"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp [][]string
	respData, err := s.client.getRessourceDataCsvURL(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return infoResp, respData, nil
}

func (s *StatsService) SeasonTTStandings(ctx context.Context, req *types.StatsSeasonTTStandingsReq) ([][]string, *http.Response, error) {
	path := "/data/stats/season_tt_standings"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp [][]string
	respData, err := s.client.getRessourceDataCsvURL(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return infoResp, respData, nil
}

func (s *StatsService) SeasonTTResults(ctx context.Context, req *types.StatsSeasonTTResultsReq) ([][]string, *http.Response, error) {
	path := "/data/stats/season_tt_results"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp [][]string
	respData, err := s.client.getRessourceDataCsvURL(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return infoResp, respData, nil
}
