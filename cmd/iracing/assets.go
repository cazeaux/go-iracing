package iracing

import (
	"context"
	"strconv"
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

func (c *IracingAssetsCache) GetCarAsset(carID int) (*types.CarAsset, bool) {
	if c.Cars == nil {
		return nil, false
	}
	carAsset, ok := (*c.Cars)[strconv.Itoa(carID)]
	if !ok {
		return nil, false
	}
	return &carAsset, true
}

func (c *IracingAssetsCache) GetTrackAsset(trackID int) (*types.TrackAsset, bool) {
	if c.Tracks == nil {
		return nil, false
	}
	trackAsset, ok := (*c.Tracks)[strconv.Itoa(trackID)]
	if !ok {
		return nil, false
	}
	return &trackAsset, true
}

func (c *IracingAssetsCache) GetSeriesAsset(seriesID int) (*types.SeriesAsset, bool) {
	if c.Series == nil {
		return nil, false
	}
	seriesAsset, ok := (*c.Series)[strconv.Itoa(seriesID)]
	if !ok {
		return nil, false
	}
	return &seriesAsset, true
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
