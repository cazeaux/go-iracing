package iracing

import (
	"context"
	"net/http"
)

type DriverStatsService struct{ client *Client }

// generic function to get driver stats by category
func (s *DriverStatsService) ByCategory(ctx context.Context, category string) ([][]string, *http.Response, error) {
	path := "/data/driver_stats_by_category/" + category
	var infoResp [][]string

	respData, err := s.client.getRessourceCSV(ctx, path, nil, &infoResp)
	if err != nil {
		return nil, respData, err
	}

	return infoResp, respData, nil
}

// specific functions for known categories

func (s *DriverStatsService) Oval(ctx context.Context) ([][]string, *http.Response, error) {
	return s.ByCategory(ctx, "oval")
}

func (s *DriverStatsService) Road(ctx context.Context) ([][]string, *http.Response, error) {
	return s.ByCategory(ctx, "road")
}

func (s *DriverStatsService) DirtOval(ctx context.Context) ([][]string, *http.Response, error) {
	return s.ByCategory(ctx, "dirt_oval")
}

func (s *DriverStatsService) DirtRoad(ctx context.Context) ([][]string, *http.Response, error) {
	return s.ByCategory(ctx, "dirt_road")
}

func (s *DriverStatsService) SportsCar(ctx context.Context) ([][]string, *http.Response, error) {
	return s.ByCategory(ctx, "sports_car")
}

func (s *DriverStatsService) FormulaCar(ctx context.Context) ([][]string, *http.Response, error) {
	return s.ByCategory(ctx, "formula_car")
}
