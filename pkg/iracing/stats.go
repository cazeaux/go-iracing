package iracing

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cazeaux/go-iracing/pkg/types"
	"github.com/google/go-querystring/query"
)

type StatsService struct{ client *Client }

func (s *StatsService) MemberRecentResult(ctx context.Context, req *types.StatsMemberRecentRacesReq) (*types.StatsMemberRecentRacesResp, *http.Response, error) {
	path := "/data/stats/member_recent_races"
	var infoResp types.StatsMemberRecentRacesResp

	params := url.Values{}
	params.Add("cust_id", strconv.Itoa(req.CustID))

	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return &infoResp, respData, nil
}

func (s *StatsService) MemberCareer(ctx context.Context, req *types.StatsMemberCareerReq) (*types.StatsMemberCareerResp, *http.Response, error) {
	path := "/data/member/member_career"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.StatsMemberCareerResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *StatsService) MemberDivision(ctx context.Context, req *types.StatsMemberDivisionReq) (*types.StatsMemberDivisionResp, *http.Response, error) {
	path := "/data/member/member_division"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.StatsMemberDivisionResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *StatsService) MemberSummary(ctx context.Context, req *types.StatsMemberSummaryReq) (*types.StatsMemberSummaryResp, *http.Response, error) {
	path := "/data/member/member_summary"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.StatsMemberSummaryResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}

func (s *StatsService) MemberYearly(ctx context.Context, req *types.StatsMemberYearlyReq) (*types.StatsMemberYearlyResp, *http.Response, error) {
	path := "/data/member/member_yearly"

	params, err := query.Values(req)
	if err != nil {
		return nil, nil, err
	}

	var infoResp types.StatsMemberYearlyResp
	respData, err := s.client.getRessourceJSON(ctx, path, params, &infoResp)
	if err != nil {
		return nil, respData, err
	}
	return &infoResp, respData, nil
}
