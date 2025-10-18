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

func (s *MemberService) Awards(ctx context.Context, req *types.MemberAwardsReq) ([]types.MemberAwardsResp, *http.Response, error) {
	path := "/data/member/awards"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp []types.MemberAwardsResp
	respData, err := s.client.getRessourceDataJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return infoResp, respData, nil
}

func (s *MemberService) AwardsInstances(ctx context.Context, req *types.MemberAwardsInstancesReq) (*types.MemberAwardsInstancesResp, *http.Response, error) {
	path := "/data/member/awards_instances"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.MemberAwardsInstancesResp
	respData, err := s.client.getRessourceDataJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *MemberService) ChartData(ctx context.Context, req *types.MemberChartDataReq) (*types.MemberChartDataResp, *http.Response, error) {
	path := "/data/member/chart_data"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.MemberChartDataResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *MemberService) Profile(ctx context.Context, req *types.MemberProfileReq) (*types.MemberProfileResp, *http.Response, error) {
	path := "/data/member/profile"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.MemberProfileResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *MemberService) ParticipationCredits(ctx context.Context, _ *types.MemberParticipationCreditsReq) ([]types.MemberParticipationCreditsResp, *http.Response, error) {
	path := "/data/member/participation_credits"

	var infoResp []types.MemberParticipationCreditsResp
	respData, err := s.client.getRessourceJSON(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return infoResp, respData, nil
}
