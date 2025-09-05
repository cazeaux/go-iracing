package types

import "time"

type StatsMemberRecentRacesReq struct {
	CustID int
}

type StatsMemberRecentRacesResp struct {
	Races []struct {
		SeasonID   int    `json:"season_id"`
		SeriesID   int    `json:"series_id"`
		SeriesName string `json:"series_name"`
		CarID      int    `json:"car_id"`
		CarClassID int    `json:"car_class_id"`
		Livery     struct {
			CarID   int    `json:"car_id"`
			Pattern int    `json:"pattern"`
			Color1  string `json:"color1"`
			Color2  string `json:"color2"`
			Color3  string `json:"color3"`
		} `json:"livery"`
		LicenseLevel     int       `json:"license_level"`
		SessionStartTime time.Time `json:"session_start_time"`
		WinnerGroupID    int       `json:"winner_group_id"`
		WinnerName       string    `json:"winner_name"`
		WinnerHelmet     struct {
			Pattern    int    `json:"pattern"`
			Color1     string `json:"color1"`
			Color2     string `json:"color2"`
			Color3     string `json:"color3"`
			FaceType   int    `json:"face_type"`
			HelmetType int    `json:"helmet_type"`
		} `json:"winner_helmet"`
		WinnerLicenseLevel int `json:"winner_license_level"`
		StartPosition      int `json:"start_position"`
		FinishPosition     int `json:"finish_position"`
		QualifyingTime     int `json:"qualifying_time"`
		Laps               int `json:"laps"`
		LapsLed            int `json:"laps_led"`
		Incidents          int `json:"incidents"`
		Points             int `json:"points"`
		StrengthOfField    int `json:"strength_of_field"`
		SubsessionID       int `json:"subsession_id"`
		OldSubLevel        int `json:"old_sub_level"`
		NewSubLevel        int `json:"new_sub_level"`
		OldiRating         int `json:"oldi_rating"`
		NewiRating         int `json:"newi_rating"`
		Track              struct {
			TrackID   int    `json:"track_id"`
			TrackName string `json:"track_name"`
		} `json:"track"`
		DropRace      bool `json:"drop_race"`
		SeasonYear    int  `json:"season_year"`
		SeasonQuarter int  `json:"season_quarter"`
		RaceWeekNum   int  `json:"race_week_num"`
	} `json:"races"`
	CustID int `json:"cust_id"`
}

// Member Career

type StatsMemberCareerReq struct {
	CustID          int  `url:"cust_id"`
	IncludeLicenses bool `url:"include_licenses,omitempty"`
}

type StatsMemberCareerResp struct {
	Stats []struct {
		CategoryID        int    `json:"category_id"`
		Category          string `json:"category"`
		Starts            int    `json:"starts"`
		Wins              int    `json:"wins"`
		Top5              int    `json:"top5"`
		Poles             int    `json:"poles"`
		AvgStartPosition  int    `json:"avg_start_position"`
		AvgFinishPosition int    `json:"avg_finish_position"`
		Laps              int    `json:"laps"`
		LapsLed           int    `json:"laps_led"`
		AvgIncidents      int    `json:"avg_incidents"`
		AvgPoints         int    `json:"avg_points"`
		WinPercentage     int    `json:"win_percentage"`
		Top5Percentage    int    `json:"top5_percentage"`
		LapsLedPercentage int    `json:"laps_led_percentage"`
		PolesPercentage   int    `json:"poles_percentage"`
	} `json:"stats"`
	CustID int `json:"cust_id"`
}

// Member Division

type StatsMemberDivisionReq struct {
	CustID    int `url:"cust_id"`
	EventType int `url:"event_type,omitempty"`
}

type StatsMemberDivisionResp struct {
	Division  int  `json:"division"`
	Projected bool `json:"projected"`
	EventType int  `json:"event_type"`
	Success   bool `json:"success"`
	SeasonID  int  `json:"season_id"`
}

// Member Summary

type StatsMemberSummaryReq struct {
	CustID int `url:"cust_id"`
}

type StatsMemberSummaryResp struct {
	ThisYear struct {
		NumOfficialSessions int `json:"num_official_sessions"`
		NumLeagueSessions   int `json:"num_league_sessions"`
		NumOfficialWins     int `json:"num_official_wins"`
		NumLeagueWins       int `json:"num_league_wins"`
	} `json:"this_year"`
	CustID int `json:"cust_id"`
}

// Member Yearly

type StatsMemberYearlyReq struct {
	CustID int `url:"cust_id"`
}

type StatsMemberYearlyResp struct {
	Stats []struct {
		CategoryID        int     `json:"category_id"`
		Category          string  `json:"category"`
		Starts            int     `json:"starts"`
		Wins              int     `json:"wins"`
		Top5              int     `json:"top5"`
		Poles             int     `json:"poles"`
		AvgStartPosition  int     `json:"avg_start_position"`
		AvgFinishPosition int     `json:"avg_finish_position"`
		Laps              int     `json:"laps"`
		LapsLed           int     `json:"laps_led"`
		AvgIncidents      float64 `json:"avg_incidents"`
		AvgPoints         int     `json:"avg_points"`
		WinPercentage     int     `json:"win_percentage"`
		Top5Percentage    float64 `json:"top5_percentage"`
		LapsLedPercentage int     `json:"laps_led_percentage"`
		Year              int     `json:"year"`
		PolesPercentage   int     `json:"poles_percentage"`
	} `json:"stats"`
	CustID int `json:"cust_id"`
}
