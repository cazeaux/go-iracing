package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
	"github.com/google/go-querystring/query"
)

type LeagueService struct{ client *Client }

func (s *LeagueService) Get(ctx context.Context, req *types.LeagueGetReq) (*types.LeagueGetResp, *http.Response, error) {
	path := "/data/league/get"
	var infoResp types.LeagueGetResp

	params, _ := query.Values(req)

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *LeagueService) CustLeagueSessions(ctx context.Context, req *types.LeagueCustLeagueSessionsReq) (*types.LeagueCustLeagueSessionsResp, *http.Response, error) {
	path := "/data/league/cust_league_sessions"
	var infoResp types.LeagueCustLeagueSessionsResp

	params, _ := query.Values(req)

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *LeagueService) Directory(ctx context.Context, req *types.LeagueDirectoryReq) (*types.LeagueDirectoryResp, *http.Response, error) {
	path := "/data/league/directory"
	var infoResp types.LeagueDirectoryResp

	params, _ := query.Values(req)

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *LeagueService) GetPointsSystems(ctx context.Context, req *types.LeagueGetPointsSystemsReq) (*types.LeagueCustLeagueSessionsResp, *http.Response, error) {
	path := "/data/league/get_points_systems"
	var infoResp types.LeagueCustLeagueSessionsResp

	params, _ := query.Values(req)

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *LeagueService) Membership(ctx context.Context, req *types.LeagueMembershipReq) (*types.LeagueMembershipResp, *http.Response, error) {
	path := "/data/league/membership"
	var infoResp types.LeagueMembershipResp

	params, _ := query.Values(req)

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *LeagueService) Seasons(ctx context.Context, req *types.LeagueSeasonsReq) (*types.LeagueSeasonsResp, *http.Response, error) {
	path := "/data/league/seasons"
	var infoResp types.LeagueSeasonsResp

	params, _ := query.Values(req)

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *LeagueService) SeasonStandings(ctx context.Context, req *types.LeagueSeasonStandingsReq) (*types.LeagueSeasonStandingsResp, *http.Response, error) {
	path := "/data/league/season_standings"
	var infoResp types.LeagueSeasonStandingsResp

	params, _ := query.Values(req)

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *LeagueService) SeasonSessions(ctx context.Context, req *types.LeagueSeasonSessionsReq) (*types.LeagueSeasonSessionsResp, *http.Response, error) {
	path := "/data/league/season_sessions"
	var infoResp types.LeagueSeasonSessionsResp

	params, _ := query.Values(req)

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}
