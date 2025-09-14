package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
	"github.com/google/go-querystring/query"
)

type MemberService struct{ client *Client }

func (s *MemberService) Info(ctx context.Context) (*types.MemberInfoResponse, *http.Response, error) {
	path := "/data/member/info"

	var infoResp types.MemberInfoResponse
	respData, err := s.client.getRessourceJSON(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *MemberService) Get(ctx context.Context, req *types.MemberGetReq) (*types.MemberGetResp, *http.Response, error) {
	path := "/data/member/get"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.MemberGetResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}
