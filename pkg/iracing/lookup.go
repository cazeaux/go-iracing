package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
	"github.com/google/go-querystring/query"
)

type LookupService struct{ client *Client }

func (s *LookupService) Licenses(ctx context.Context, _ *types.LookupLicensesReq) ([]types.LookupLicensesResp, *http.Response, error) {
	path := "/data/lookup/licenses"
	var infoResp []types.LookupLicensesResp

	respData, err := s.client.getRessourceJSON(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}

func (s *LookupService) Countries(ctx context.Context, _ *types.LookupCountriesReq) ([]types.LookupCountriesResp, *http.Response, error) {
	path := "/data/lookup/countries"
	var infoResp []types.LookupCountriesResp

	respData, err := s.client.getRessourceJSON(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}

func (s *LookupService) Flairs(ctx context.Context, _ *types.LookupFlairsReq) ([]types.LookupFlairsResp, *http.Response, error) {
	path := "/data/lookup/flairs"
	var infoResp []types.LookupFlairsResp

	respData, err := s.client.getRessourceJSON(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}

func (s *LookupService) Drivers(ctx context.Context, req *types.LookupDriversReq) ([]types.LookupDriversResp, *http.Response, error) {
	path := "/data/lookup/drivers"
	var infoResp []types.LookupDriversResp

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}
