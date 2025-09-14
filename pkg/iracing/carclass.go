package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
)

type CarClassService struct{ client *Client }

func (s *CarClassService) Get(ctx context.Context, _ *types.CarClassGetReq) ([]types.CarClassGetResp, *http.Response, error) {
	path := "/data/carclass/get"
	var infoResp []types.CarClassGetResp

	respData, err := s.client.getRessourceJSON(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}
