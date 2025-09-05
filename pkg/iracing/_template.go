package iracing

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// Track représente un circuit iRacing (exemple minimal).
type TemplateReq struct {
}

type TemplateResp struct {
}

type TemplateService struct{ client *Client }

func (s *TemplateService) Get(ctx context.Context, req *TemplateReq) (*TemplateResp, *http.Response, error) {
	path := "/data/stats/member_recent_races"
	var infoResp TemplateResp

	params := url.Values{}
	params.Add("cust_id", strconv.Itoa(req.CustID))

	respData, err := s.client.getRessource(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}
