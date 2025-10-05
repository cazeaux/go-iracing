package iracing

import (
	"context"
	"time"

	"github.com/cazeaux/go-iracing/pkg/iracing"
	"github.com/cazeaux/go-iracing/pkg/types"
)

type IracingAssetsCache struct {
	Tracks      *types.TrackAssetsResp
	Cars        *types.CarsAssetsResp
	Series      *types.SeriesAssetsResp
	LastRefresh time.Time
	Expiry      time.Duration
}

func (c *IracingAssetsCache) IsExpired() bool {
	if c.LastRefresh.IsZero() {
		return true
	}
	return time.Since(c.LastRefresh) > c.Expiry
}

func NewIracingAssetsCache(expiry time.Duration) *IracingAssetsCache {
	return &IracingAssetsCache{
		Expiry: expiry,
	}
}

func (c *IracingAssetsCache) RefreshCache(ctx context.Context, client *iracing.Client) error {
	if !c.IsExpired() {
		return nil
	}

	var err error
	trackReq := &types.TrackAssetsReq{}
	c.Tracks, _, err = client.Track.Assets(ctx, trackReq)
	if err != nil {
		return err
	}

	carReq := &types.CarsAssetsReq{}
	c.Cars, _, err = client.Car.Assets(ctx, carReq)
	if err != nil {
		return err
	}

	seriesReq := &types.SeriesAssetsReq{}
	c.Series, _, err = client.Series.Assets(ctx, seriesReq)
	if err != nil {
		return err
	}

	c.LastRefresh = time.Now()
	return nil
}
