package iracing

import "time"

type RaceResult struct {
	BestLapTime      int //
	CarID            int //
	CarLogo          string //
	CarName          string //
	ChamPoints       int    //
	FieldSize        int    //
	Incidents        int    //
	IsMultiClass     bool //
	MultiClass       *MultiClass
	NewIRating       int //
	NewSRating       int //
	OldIRating       int //
	OldSRating       int //
	PilotDisplayName string //
	QualiLapTime     int //
	QualyPos         int //
	QualyRain        bool //
	RacePos          int //
	RaceRain         bool //
	SeriesID         int //
	SeriesLogo       string //
	SeriesName       string //
	SOF              int    //
	Split            int    //
	Splits           int    //
	StartTime        time.Time
	SubsessionID     int    //
	TrackID          int    //
	TrackName        string //
}

type MultiClass struct {
	FieldSize int
	SOF       int
	RacePos   int
	QualyPos  int
}
