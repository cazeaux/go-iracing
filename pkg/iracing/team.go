package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
	"github.com/google/go-querystring/query"
)

type TeamService struct{ client *Client }

func (s *TeamService) Get(ctx context.Context, req *types.TeamGetReq) (*types.TeamGetReq, *http.Response, error) {
	path := "/data/team/get"
	var infoResp types.TeamGetReq

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *TeamService) Membership(ctx context.Context, _ *types.TeamMembershipReq) ([]types.TeamMembershipResp, *http.Response, error) {
	path := "/data/team/get"
	var infoResp []types.TeamMembershipResp

	respData, err := s.client.getRessourceJSON(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}
