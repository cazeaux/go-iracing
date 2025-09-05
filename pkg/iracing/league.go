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

	respData, err := s.client.getRessource(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}
