package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
	"github.com/google/go-querystring/query"
)

type ResultsService struct{ client *Client }

func (s *ResultsService) SearchSeries(ctx context.Context, req *types.ResultsSearchSeriesReq) ([]types.ResultsSearchSeriesResp, *http.Response, error) {
	path := "/data/results/search_series"
	var infoResp []types.ResultsSearchSeriesResp

	params, _ := query.Values(req)

	respData, err := GetRessourceChunks(ctx, s.client, path, params, req.ChunkReq.Rows, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}

func (s *ResultsService) Get(ctx context.Context, req *types.ResultsGetReq) (*types.ResultsGetResp, *http.Response, error) {
	path := "/data/results/get"
	var infoResp types.ResultsGetResp

	params, _ := query.Values(req)

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *ResultsService) SeasonResults(ctx context.Context, req *types.ResultsSeasonResultsReq) (*types.ResultsSeasonResultsResp, *http.Response, error) {
	path := "/data/results/season_results"
	var infoResp types.ResultsSeasonResultsResp

	params, _ := query.Values(req)

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}
