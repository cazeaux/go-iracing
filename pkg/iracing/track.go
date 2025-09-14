package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
)

type TrackService struct {
	client *Client
}

func (s *TrackService) Get(ctx context.Context, _ *types.TrackGetReq) ([]types.TrackGetResp, *http.Response, error) {
	path := "/data/track/get"
	var infoResp []types.TrackGetResp

	respData, err := s.client.getRessourceJSON(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}
