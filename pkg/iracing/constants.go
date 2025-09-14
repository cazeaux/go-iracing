package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
)

type ConstantsService struct{ client *Client }

func (s *ConstantsService) Divisions(ctx context.Context) ([]types.ConstantsDivisionsResp, *http.Response, error) {
	path := "/data/constants/divisions"
	var infoResp []types.ConstantsDivisionsResp

	respData, err := s.client.get(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}

func (s *ConstantsService) Categories(ctx context.Context) ([]types.ConstantsCategoriesResp, *http.Response, error) {
	path := "/data/constants/categories"
	var infoResp []types.ConstantsCategoriesResp

	respData, err := s.client.get(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}

func (s *ConstantsService) EventTypes(ctx context.Context) ([]types.ConstantsEventTypesResp, *http.Response, error) {
	path := "/data/constants/eventtypes"
	var infoResp []types.ConstantsEventTypesResp

	respData, err := s.client.get(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}
