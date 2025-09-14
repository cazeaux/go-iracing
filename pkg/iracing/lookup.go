package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
)

type LookupService struct{ client *Client }

func (s *LookupService) Licenses(ctx context.Context, req *types.LookupLicensesReq) ([]types.LookupLicensesResp, *http.Response, error) {
	path := "/data/lookup/licenses"
	var infoResp []types.LookupLicensesResp

	respData, err := s.client.getRessourceJSON(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}
