package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
	"github.com/google/go-querystring/query"
)

type HostedService struct{ client *Client }

func (s *HostedService) CombinedSessions(ctx context.Context, req types.HostedCombinedSessionsReq) (*types.HostedCombinedSessionsResp, *http.Response, error) {
	path := "/data/hosted/combined_sessions"
	var infoResp types.HostedCombinedSessionsResp

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

func (s *HostedService) Sessions(ctx context.Context, req types.HostedSessionsReq) (*types.HostedSessionsResp, *http.Response, error) {
	path := "/data/hosted/sessions"
	var infoResp types.HostedSessionsResp

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
