package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
	"github.com/google/go-querystring/query"
)

type SeriesService struct {
	client *Client
}

func (s *SeriesService) Seasons(ctx context.Context, _ *types.SeriesSeasonsReq) ([]types.SeriesSeasonsResp, *http.Response, error) {
	path := "/data/series/seasons"
	var infoResp []types.SeriesSeasonsResp

	respData, err := s.client.getRessourceJSON(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}

func (s *SeriesService) StatsSeries(ctx context.Context, req *types.SeriesStatsSeriesReq) ([]types.SeriesStatsSeriesResp, *http.Response, error) {
	path := "/data/series/stats_series"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp []types.SeriesStatsSeriesResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}

func (s *SeriesService) Assets(ctx context.Context, _ *types.SeriesAssetsReq) (*types.SeriesAssetsResp, *http.Response, error) {
	path := "/data/series/assets"

	var infoResp types.SeriesAssetsResp
	respData, err := s.client.getRessourceJSON(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *SeriesService) Get(ctx context.Context, _ *types.SeriesGetReq) ([]types.SeriesGetResp, *http.Response, error) {
	path := "/data/series/get"

	var infoResp []types.SeriesGetResp
	respData, err := s.client.getRessourceJSON(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}

func (s *SeriesService) PastSeasons(ctx context.Context, req *types.SeriesPastSeasonsReq) (*types.SeriesPastSeasonsResp, *http.Response, error) {
	path := "/data/series/past_seasons"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.SeriesPastSeasonsResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *SeriesService) SeasonList(ctx context.Context, req *types.SeriesSeasonListReq) (*types.SeriesSeasonListResp, *http.Response, error) {
	path := "/data/series/season_list"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.SeriesSeasonListResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *SeriesService) SeasonSchedule(ctx context.Context, req *types.SeriesSeasonScheduleReq) (*types.SeriesSeasonScheduleResp, *http.Response, error) {
	path := "/data/series/season_schedule"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.SeriesSeasonScheduleResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}
