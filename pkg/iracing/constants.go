package iracing

import (
	"context"
	"net/http"

	"github.com/cazeaux/go-iracing/pkg/types"
)

type ConstantsService struct{ client *Client }

func (s *ConstantsService) Divisions(ctx context.Context, req *types.ConstantsDivisionsReq) ([]types.ConstantsDivisionsResp, *http.Response, error) {
	path := "/data/stats/member_recent_races"
	var infoResp []types.ConstantsDivisionsResp

	respData, err := s.client.getRessource(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}
