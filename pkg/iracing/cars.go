package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
)

type CarService struct{ client *Client }

func (s *CarService) Get(ctx context.Context, _ *types.CarsGetReq) ([]types.CarsGetResp, *http.Response, error) {
	path := "/data/car/get"
	var infoResp []types.CarsGetResp

	respData, err := s.client.getRessource(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}

func (s *CarService) Assets(ctx context.Context, _ *types.CarsAssetsReq) (*types.CarsAssetsResp, *http.Response, error) {
	path := "/data/car/assets"
	var infoResp types.CarsAssetsResp

	respData, err := s.client.getRessource(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}
