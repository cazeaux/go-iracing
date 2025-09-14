package iracing

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

type TemplateService struct{ client *Client }

func (s *TemplateService) Get(ctx context.Context, req *types.TemplateReq) (*types.TemplateResp, *http.Response, error) {
	path := "/data/stats/member_recent_races"
	var infoResp types.TemplateResp

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
