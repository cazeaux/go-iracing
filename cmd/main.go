package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cazeaux/go-iracing/pkg/iracing"
	"github.com/cazeaux/go-iracing/pkg/types"
)

func main() {
	ctx := context.Background()
	email := os.Getenv("IR_EMAIL")
	password := os.Getenv("IR_PASSWORD")
	clientID := os.Getenv("IR_CLIENT_ID")
	clientSecret := os.Getenv("IR_CLIENT_SECRET")
	client, err := iracing.NewClient(
		// iracing.WithAuthenticator(iracing.NewIrLegacyAuth(email, password)),
		iracing.WithAuthenticator(iracing.NewIrOAuth(email, password, clientID, clientSecret)),
	)
	if err != nil {
		log.Fatalf("new client: %v", err)
	}

	// getMemberInfo(client, ctx)
	// getMemberRecentResult(client, ctx, 394410)
	// getCars(client, ctx)
	// getTracks(client, ctx)
	// getResultsSearchSeries(client, ctx, 394410)
	// lookupLicenses(client, ctx)
	// getResultsGet(client, ctx, 78354169)
	// getSeriesSeasons(client, ctx)
	// getStatsSeries(client, ctx)
	// getDriverStats(client, ctx)
	// getHostedCombinedSessions(client, ctx)
	// getHostedSessions(client, ctx)
	// getMemberAwards(client, ctx)
	// getSeasonQualifyResults(client, ctx)
	getSeasonDriverStandings(client, ctx)

}

func getMemberInfo(c *iracing.Client, ctx context.Context) {
	info, resp, err := c.Member.Info(ctx)
	if err != nil {
		log.Fatalf("info: %v", err)
	}
	fmt.Printf("Info %v: %v\n", resp.StatusCode, info)
}

func getMemberRecentResult(c *iracing.Client, ctx context.Context, custID int) {
	req := &types.StatsMemberRecentRacesReq{CustID: custID}
	recentResult, resp, err := c.Stats.MemberRecentResult(ctx, req)
	if err != nil {
		log.Fatalf("info: %v", err)
	}
	fmt.Printf("Info %v: %v\n", resp.StatusCode, recentResult)
}

func getCars(c *iracing.Client, ctx context.Context) {
	cars, resp, err := c.Car.Assets(ctx, nil)
	if err != nil {
		log.Fatalf("list cars failed: %v", err)
	}
	log.Printf("cars: %v (status %d)", cars, resp.StatusCode)
}

func getTracks(c *iracing.Client, ctx context.Context) {
	tracks, resp, err := c.Track.Get(ctx, nil)
	if err != nil {
		log.Fatalf("list tracks failed: %v", err)
	}
	log.Printf("tracks: %d (status %d)", len(tracks), resp.StatusCode)
}

func lookupLicenses(c *iracing.Client, ctx context.Context) {
	data, resp, err := c.LookupService.Licenses(ctx, nil)
	if err != nil {
		log.Fatalf("lookup licenses failed: %v", err)
	}
	log.Printf("%v %v", resp.StatusCode, data)
}

func getResultsSearchSeries(c *iracing.Client, ctx context.Context, custID int) {
	layout := "2006-01-02"
	value := "2025-06-15"

	date, _ := time.Parse(layout, value)

	req := &types.ResultsSearchSeriesReq{
		CustID:          custID,
		SeasonYear:      2025,
		SeasonQuarter:   3,
		StartRangeBegin: date,
		ChunkReq: types.ChunkReq{
			Rows: 100,
		},
	}
	results, _, err := c.Results.SearchSeries(ctx, req)
	if err != nil {
		log.Fatalf("list results failed: %v", err)
	}
	log.Printf("results: %d", len(results))
}

func getResultsGet(c *iracing.Client, ctx context.Context, subsessionID int) {
	req := &types.ResultsGetReq{
		SubsessionID: subsessionID,
	}
	results, _, err := c.Results.Get(ctx, req)
	if err != nil {
		log.Fatalf("list results failed: %v", err)
	}
	log.Printf("results: %v", results)
}

func getSeriesSeasons(c *iracing.Client, ctx context.Context) {
	req := &types.SeriesSeasonsReq{
		IncludeSeries: true,
	}
	results, _, err := c.SeriesService.Seasons(ctx, req)
	if err != nil {
		log.Fatalf("list results failed: %v", err)
	}
	log.Printf("results: %v", results)
}

func getStatsSeries(c *iracing.Client, ctx context.Context) {
	req := &types.SeriesStatsSeriesReq{
		Official: true,
	}
	results, _, err := c.SeriesService.StatsSeries(ctx, req)
	if err != nil {
		log.Fatalf("list results failed: %v", err)
	}
	log.Printf("results: %v", results)
}

func getDriverStats(c *iracing.Client, ctx context.Context) {

	results, _, err := c.DriverStatsService.Oval(ctx)
	if err != nil {
		log.Fatalf("list results failed: %v", err)
	}
	log.Printf("results: %v", results[:10])
}

func getHostedCombinedSessions(c *iracing.Client, ctx context.Context) {
	req := types.HostedCombinedSessionsReq{}
	results, _, err := c.HostedService.CombinedSessions(ctx, req)
	if err != nil {
		log.Fatalf("list results failed: %v", err)
	}
	log.Printf("results: %v", results)
}

func getHostedSessions(c *iracing.Client, ctx context.Context) {
	req := types.HostedSessionsReq{}
	results, _, err := c.HostedService.Sessions(ctx, req)
	if err != nil {
		log.Fatalf("list results failed: %v", err)
	}
	log.Printf("results: %v", results)
}

func getMemberAwards(c *iracing.Client, ctx context.Context) {
	req := types.MemberAwardsReq{CustID: 394410}
	results, _, err := c.MemberService.Awards(ctx, &req)
	if err != nil {
		log.Fatalf("list results failed: %v", err)
	}
	log.Printf("results: %v", results)
}

func getSeasonQualifyResults(c *iracing.Client, ctx context.Context) {
	req := types.StatsSeasonQualifyResultsReq{CarClassID: 2708, SeasonID: 5420, RaceWeekNum: 1}
	results, _, err := c.Stats.SeasonQualifyResults(ctx, &req)
	if err != nil {
		log.Fatalf("list results failed: %v", err)
	}
	log.Printf("results: %v", results)
}

func getSeasonDriverStandings(c *iracing.Client, ctx context.Context) {
	req := types.StatsSeasonDriverStandingsReq{CarClassID: 2708, SeasonID: 5420, RaceWeekNum: 1}
	results, _, err := c.Stats.SeasonDriverStandings(ctx, &req)
	if err != nil {
		log.Fatalf("list results failed: %v", err)
	}
	log.Printf("results: %v", results)
}
