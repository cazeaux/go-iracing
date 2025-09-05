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

	respData, err := s.client.getRessource(ctx, path, nil, &infoResp)
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
	respData, err := s.client.getRessource(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}
