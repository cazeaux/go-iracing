package types

import "time"

// Search

type ResultsSearchSeriesReq struct {
	ChunkReq         ChunkReq
	CustID           int       `url:"cust_id"`
	StartRangeBegin  time.Time `url:"start_range_begin"`
	StartRangeEnd    time.Time `url:"start_range_end,omitempty"`
	FinishRangeBegin time.Time `url:"finish_range_begin,omitempty"`
	FinishRangeEnd   time.Time `url:"finish_range_end,omitempty"`
	SeriesID         int       `url:"series_id,omitempty"`
	TeamID           int       `url:"team_id,omitempty"`
	SeasonYear       int       `url:"season_year,omitempty"`
	SeasonQuarter    int       `url:"season_quarter,omitempty"`
	RaceWeekNum      int       `url:"race_week_num,omitempty"`
	OfficialOnly     *bool     `url:"official_only,omitempty"`
	EventTypes       []int     `url:"event_types,omitempty"`
	CategoryIDs      []int     `url:"category_ids,omitempty"`
}

type ResultsSearchSeriesResp struct {
	SessionID               int       `json:"session_id"`
	SubsessionID            int       `json:"subsession_id"`
	StartTime               time.Time `json:"start_time"`
	EndTime                 time.Time `json:"end_time"`
	LicenseCategoryID       int       `json:"license_category_id"`
	LicenseCategory         string    `json:"license_category"`
	NumDrivers              int       `json:"num_drivers"`
	NumCautions             int       `json:"num_cautions"`
	NumCautionLaps          int       `json:"num_caution_laps"`
	NumLeadChanges          int       `json:"num_lead_changes"`
	EventAverageLap         int       `json:"event_average_lap"`
	EventBestLapTime        int       `json:"event_best_lap_time"`
	EventLapsComplete       int       `json:"event_laps_complete"`
	DriverChanges           bool      `json:"driver_changes"`
	WinnerGroupID           int       `json:"winner_group_id"`
	WinnerName              string    `json:"winner_name"`
	WinnerAi                bool      `json:"winner_ai"`
	CustID                  int       `json:"cust_id"`
	StartingPosition        int       `json:"starting_position"`
	FinishPosition          int       `json:"finish_position"`
	StartingPositionInClass int       `json:"starting_position_in_class"`
	FinishPositionInClass   int       `json:"finish_position_in_class"`
	LapsComplete            int       `json:"laps_complete"`
	LapsLed                 int       `json:"laps_led"`
	Incidents               int       `json:"incidents"`
	CarClassID              int       `json:"car_class_id"`
	CarID                   int       `json:"car_id"`
	CarClassName            string    `json:"car_class_name"`
	CarClassShortName       string    `json:"car_class_short_name"`
	CarName                 string    `json:"car_name"`
	CarNameAbbreviated      string    `json:"car_name_abbreviated"`
	Track                   struct {
		ConfigName string `json:"config_name"`
		TrackID    int    `json:"track_id"`
		TrackName  string `json:"track_name"`
	} `json:"track"`
	OfficialSession        bool   `json:"official_session"`
	SeasonID               int    `json:"season_id"`
	SeasonYear             int    `json:"season_year"`
	SeasonQuarter          int    `json:"season_quarter"`
	EventType              int    `json:"event_type"`
	EventTypeName          string `json:"event_type_name"`
	SeriesID               int    `json:"series_id"`
	SeriesName             string `json:"series_name"`
	SeriesShortName        string `json:"series_short_name"`
	RaceWeekNum            int    `json:"race_week_num"`
	EventStrengthOfField   int    `json:"event_strength_of_field"`
	ChampPoints            int    `json:"champ_points"`
	DropRace               bool   `json:"drop_race"`
	SeasonLicenseGroup     int    `json:"season_license_group"`
	SeasonLicenseGroupName string `json:"season_license_group_name"`
}

// Get

type ResultsGetReq struct {
	SubsessionID    int  `url:"subsession_id"`
	IncludeLicenses bool `url:"include_licenses,omitempty"`
}

type ResultsGetResp struct {
	SubsessionID            int                     `json:"subsession_id"`
	AllowedLicenses         []CommonAllowedLicenses `json:"allowed_licenses"`
	AssociatedSubsessionIds []int                   `json:"associated_subsession_ids"`
	CanProtest              bool                    `json:"can_protest"`
	CarClasses              []struct {
		CarClassID      int    `json:"car_class_id"`
		ShortName       string `json:"short_name"`
		Name            string `json:"name"`
		StrengthOfField int    `json:"strength_of_field"`
		NumEntries      int    `json:"num_entries"`
		CarsInClass     []struct {
			CarID int `json:"car_id"`
		} `json:"cars_in_class"`
	} `json:"car_classes"`
	CautionType           int       `json:"caution_type"`
	CooldownMinutes       int       `json:"cooldown_minutes"`
	CornersPerLap         int       `json:"corners_per_lap"`
	DamageModel           int       `json:"damage_model"`
	DriverChangeParam1    int       `json:"driver_change_param1"`
	DriverChangeParam2    int       `json:"driver_change_param2"`
	DriverChangeRule      int       `json:"driver_change_rule"`
	DriverChanges         bool      `json:"driver_changes"`
	EndTime               time.Time `json:"end_time"`
	EventAverageLap       int       `json:"event_average_lap"`
	EventBestLapTime      int       `json:"event_best_lap_time"`
	EventLapsComplete     int       `json:"event_laps_complete"`
	EventStrengthOfField  int       `json:"event_strength_of_field"`
	EventType             int       `json:"event_type"`
	EventTypeName         string    `json:"event_type_name"`
	HeatInfoID            int       `json:"heat_info_id"`
	LicenseCategory       string    `json:"license_category"`
	LicenseCategoryID     int       `json:"license_category_id"`
	LimitMinutes          int       `json:"limit_minutes"`
	MaxTeamDrivers        int       `json:"max_team_drivers"`
	MaxWeeks              int       `json:"max_weeks"`
	MinTeamDrivers        int       `json:"min_team_drivers"`
	NumCautionLaps        int       `json:"num_caution_laps"`
	NumCautions           int       `json:"num_cautions"`
	NumDrivers            int       `json:"num_drivers"`
	NumLapsForQualAverage int       `json:"num_laps_for_qual_average"`
	NumLapsForSoloAverage int       `json:"num_laps_for_solo_average"`
	NumLeadChanges        int       `json:"num_lead_changes"`
	OfficialSession       bool      `json:"official_session"`
	PointsType            string    `json:"points_type"`
	PrivateSessionID      int       `json:"private_session_id"`
	RaceSummary           struct {
		SubsessionID         int    `json:"subsession_id"`
		AverageLap           int    `json:"average_lap"`
		LapsComplete         int    `json:"laps_complete"`
		NumCautions          int    `json:"num_cautions"`
		NumCautionLaps       int    `json:"num_caution_laps"`
		NumLeadChanges       int    `json:"num_lead_changes"`
		FieldStrength        int    `json:"field_strength"`
		NumOptLaps           int    `json:"num_opt_laps"`
		HasOptPath           bool   `json:"has_opt_path"`
		SpecialEventType     int    `json:"special_event_type"`
		SpecialEventTypeText string `json:"special_event_type_text"`
	} `json:"race_summary"`
	RaceWeekNum       int    `json:"race_week_num"`
	ResultsRestricted bool   `json:"results_restricted"`
	SeasonID          int    `json:"season_id"`
	SeasonName        string `json:"season_name"`
	SeasonQuarter     int    `json:"season_quarter"`
	SeasonShortName   string `json:"season_short_name"`
	SeasonYear        int    `json:"season_year"`
	SeriesID          int    `json:"series_id"`
	SeriesLogo        string `json:"series_logo"`
	SeriesName        string `json:"series_name"`
	SeriesShortName   string `json:"series_short_name"`
	SessionID         int    `json:"session_id"`
	SessionResults    []struct {
		SimsessionNumber   int                              `json:"simsession_number"`
		SimsessionName     string                           `json:"simsession_name"`
		SimsessionType     int                              `json:"simsession_type"`
		SimsessionTypeName string                           `json:"simsession_type_name"`
		SimsessionSubtype  int                              `json:"simsession_subtype"`
		WeatherResult      SessionWeatherResult             `json:"weather_result"`
		Results            []SubsessionSessionResultsPilots `json:"results"`
	} `json:"session_results"`
	SessionSplits    []SubsessionSplit `json:"session_splits"`
	SpecialEventType int               `json:"special_event_type"`
	StartTime        time.Time         `json:"start_time"`
	Track            CommonTrack       `json:"track"`
	TrackState       CommonTrackState  `json:"track_state"`
	Weather          CommonWeather     `json:"weather"`
}

type SessionWeatherResult struct {
	AvgSkies                 int     `json:"avg_skies"`
	AvgCloudCoverPct         float64 `json:"avg_cloud_cover_pct"`
	MinCloudCoverPct         float64 `json:"min_cloud_cover_pct"`
	MaxCloudCoverPct         float64 `json:"max_cloud_cover_pct"`
	TempUnits                int     `json:"temp_units"`
	AvgTemp                  float64 `json:"avg_temp"`
	MinTemp                  float64 `json:"min_temp"`
	MaxTemp                  float64 `json:"max_temp"`
	AvgRelHumidity           float64 `json:"avg_rel_humidity"`
	WindUnits                int     `json:"wind_units"`
	AvgWindSpeed             float64 `json:"avg_wind_speed"`
	MinWindSpeed             float64 `json:"min_wind_speed"`
	MaxWindSpeed             float64 `json:"max_wind_speed"`
	AvgWindDir               int     `json:"avg_wind_dir"`
	MaxFog                   int     `json:"max_fog"`
	FogTimePct               float64 `json:"fog_time_pct"`
	PrecipTimePct            float64 `json:"precip_time_pct"`
	PrecipMm                 float64 `json:"precip_mm"`
	PrecipMm2HrBeforeSession float64 `json:"precip_mm2hr_before_session"`
	SimulatedStartTime       string  `json:"simulated_start_time"`
}
type SubsessionSplit struct {
	SubsessionID         int `json:"subsession_id"`
	EventStrengthOfField int `json:"event_strength_of_field"`
}

type SubsessionSessionResultsPilots struct {
	CustID                int          `json:"cust_id"`
	DisplayName           string       `json:"display_name"`
	AggregateChampPoints  int          `json:"aggregate_champ_points"`
	Ai                    bool         `json:"ai"`
	AverageLap            int          `json:"average_lap"`
	BestLapNum            int          `json:"best_lap_num"`
	BestLapTime           int          `json:"best_lap_time"`
	BestNlapsNum          int          `json:"best_nlaps_num"`
	BestNlapsTime         int          `json:"best_nlaps_time"`
	BestQualLapAt         time.Time    `json:"best_qual_lap_at"`
	BestQualLapNum        int          `json:"best_qual_lap_num"`
	BestQualLapTime       int          `json:"best_qual_lap_time"`
	CarClassID            int          `json:"car_class_id"`
	CarClassName          string       `json:"car_class_name"`
	CarClassShortName     string       `json:"car_class_short_name"`
	CarID                 int          `json:"car_id"`
	CarName               string       `json:"car_name"`
	Carcfg                int          `json:"carcfg"`
	ChampPoints           int          `json:"champ_points"`
	ClassInterval         int          `json:"class_interval"`
	CountryCode           string       `json:"country_code"`
	Division              int          `json:"division"`
	DivisionName          string       `json:"division_name"`
	DropRace              bool         `json:"drop_race"`
	FinishPosition        int          `json:"finish_position"`
	FinishPositionInClass int          `json:"finish_position_in_class"`
	FlairID               int          `json:"flair_id"`
	FlairName             string       `json:"flair_name"`
	FlairShortname        string       `json:"flair_shortname"`
	Friend                bool         `json:"friend"`
	Helmet                CommonHelmet `json:"helmet"`
	Incidents             int          `json:"incidents"`
	Interval              int          `json:"interval"`
	LapsComplete          int          `json:"laps_complete"`
	LapsLead              int          `json:"laps_lead"`
	LeagueAggPoints       int          `json:"league_agg_points"`
	LeaguePoints          int          `json:"league_points"`
	LicenseChangeOval     int          `json:"license_change_oval"`
	LicenseChangeRoad     int          `json:"license_change_road"`
	Livery                struct {
		CarID        int    `json:"car_id"`
		Pattern      int    `json:"pattern"`
		Color1       string `json:"color1"`
		Color2       string `json:"color2"`
		Color3       string `json:"color3"`
		NumberFont   int    `json:"number_font"`
		NumberColor1 string `json:"number_color1"`
		NumberColor2 string `json:"number_color2"`
		NumberColor3 string `json:"number_color3"`
		NumberSlant  int    `json:"number_slant"`
		Sponsor1     int    `json:"sponsor1"`
		Sponsor2     int    `json:"sponsor2"`
		CarNumber    string `json:"car_number"`
		WheelColor   any    `json:"wheel_color"`
		RimType      int    `json:"rim_type"`
	} `json:"livery"`
	MaxPctFuelFill          float64 `json:"max_pct_fuel_fill"`
	NewCpi                  float64 `json:"new_cpi"`
	NewLicenseLevel         int     `json:"new_license_level"`
	NewSubLevel             int     `json:"new_sub_level"`
	NewTtrating             int     `json:"new_ttrating"`
	NewiRating              int     `json:"newi_rating"`
	OldCpi                  float64 `json:"old_cpi"`
	OldLicenseLevel         int     `json:"old_license_level"`
	OldSubLevel             int     `json:"old_sub_level"`
	OldTtrating             int     `json:"old_ttrating"`
	OldiRating              int     `json:"oldi_rating"`
	OptLapsComplete         int     `json:"opt_laps_complete"`
	Position                int     `json:"position"`
	QualLapTime             int     `json:"qual_lap_time"`
	ReasonOut               string  `json:"reason_out"`
	ReasonOutID             int     `json:"reason_out_id"`
	StartingPosition        int     `json:"starting_position"`
	StartingPositionInClass int     `json:"starting_position_in_class"`
	Suit                    struct {
		Pattern int    `json:"pattern"`
		Color1  string `json:"color1"`
		Color2  string `json:"color2"`
		Color3  string `json:"color3"`
	} `json:"suit"`
	Watched         bool `json:"watched"`
	WeightPenaltyKg int  `json:"weight_penalty_kg"`
}

// SeasonResults

type ResultsSeasonResultsReq struct {
	SeasonID    int `url:"season_id"`
	RaceWeekNum any `url:"race_week_num,omitempty"`
	EventType   int `url:"event_type,omitempty"`
}

type ResultsSeasonResultsResp struct {
	Success     bool `json:"success"`
	SeasonID    int  `json:"season_id"`
	RaceWeekNum any  `json:"race_week_num"`
	EventType   int  `json:"event_type"`
	ResultsList []struct {
		SessionID    int `json:"session_id"`
		SubsessionID int `json:"subsession_id"`
		RaceWeekNum  int `json:"race_week_num"`
		CarClasses   []struct {
			CarClassID      int    `json:"car_class_id"`
			ShortName       string `json:"short_name"`
			Name            string `json:"name"`
			NumEntries      int    `json:"num_entries"`
			StrengthOfField int    `json:"strength_of_field"`
		} `json:"car_classes"`
		DriverChanges        bool   `json:"driver_changes"`
		EventBestLapTime     int    `json:"event_best_lap_time"`
		EventStrengthOfField int    `json:"event_strength_of_field"`
		EventType            int    `json:"event_type"`
		EventTypeName        string `json:"event_type_name"`
		Farm                 struct {
			FarmID      int    `json:"farm_id"`
			DisplayName string `json:"display_name"`
			ImagePath   string `json:"image_path"`
			Displayed   bool   `json:"displayed"`
		} `json:"farm"`
		NumCautionLaps  int       `json:"num_caution_laps"`
		NumCautions     int       `json:"num_cautions"`
		NumDrivers      int       `json:"num_drivers"`
		NumLeadChanges  int       `json:"num_lead_changes"`
		OfficialSession bool      `json:"official_session"`
		StartTime       time.Time `json:"start_time"`
		Track           struct {
			ConfigName string `json:"config_name"`
			TrackID    int    `json:"track_id"`
			TrackName  string `json:"track_name"`
		} `json:"track"`
		WinnerHelmet struct {
			Pattern    int    `json:"pattern"`
			Color1     string `json:"color1"`
			Color2     string `json:"color2"`
			Color3     string `json:"color3"`
			FaceType   int    `json:"face_type"`
			HelmetType int    `json:"helmet_type"`
		} `json:"winner_helmet"`
		WinnerID           int    `json:"winner_id"`
		WinnerLicenseLevel int    `json:"winner_license_level"`
		WinnerName         string `json:"winner_name"`
	} `json:"results_list"`
}

// search_hosted

type ResultsSearchHostedReq struct {
	ChunkReq         ChunkReq
	CustID           int       `url:"cust_id"`
	StartRangeBegin  time.Time `url:"start_range_begin"`
	StartRangeEnd    time.Time `url:"start_range_end,omitempty"`
	FinishRangeBegin time.Time `url:"finish_range_begin,omitempty"`
	FinishRangeEnd   time.Time `url:"finish_range_end,omitempty"`
	TeamID           int       `url:"team_id,omitempty"`
	HostCustID       int       `url:"host_cust_id,omitempty"`
	SessionName      string    `url:"session_name,omitempty"`
	LeagueID         int       `url:"league_id,omitempty"`
	LeagueSeasonID   int       `url:"league_season_id,omitempty"`
	CarID            int       `url:"car_id,omitempty"`
	TrackID          int       `url:"track_id,omitempty"`
	CategoryIDs      []int     `url:"category_ids,omitempty"`
}

type ResultsSearchHostedResp []struct {
	SessionID               int         `json:"session_id"`
	SubsessionID            int         `json:"subsession_id"`
	StartTime               time.Time   `json:"start_time"`
	EndTime                 time.Time   `json:"end_time"`
	LicenseCategoryID       int         `json:"license_category_id"`
	LicenseCategory         string      `json:"license_category"`
	NumDrivers              int         `json:"num_drivers"`
	NumCautions             int         `json:"num_cautions"`
	NumCautionLaps          int         `json:"num_caution_laps"`
	NumLeadChanges          int         `json:"num_lead_changes"`
	EventAverageLap         int         `json:"event_average_lap"`
	EventBestLapTime        int         `json:"event_best_lap_time"`
	EventLapsComplete       int         `json:"event_laps_complete"`
	DriverChanges           bool        `json:"driver_changes"`
	WinnerGroupID           int         `json:"winner_group_id"`
	WinnerName              string      `json:"winner_name"`
	WinnerAi                bool        `json:"winner_ai"`
	CustID                  int         `json:"cust_id"`
	StartingPosition        int         `json:"starting_position"`
	FinishPosition          int         `json:"finish_position"`
	StartingPositionInClass int         `json:"starting_position_in_class"`
	FinishPositionInClass   int         `json:"finish_position_in_class"`
	LapsComplete            int         `json:"laps_complete"`
	LapsLed                 int         `json:"laps_led"`
	Incidents               int         `json:"incidents"`
	Track                   CommonTrack `json:"track"`
	PrivateSessionID        int         `json:"private_session_id"`
	SessionName             string      `json:"session_name"`
	LeagueID                int         `json:"league_id"`
	LeagueSeasonID          int         `json:"league_season_id"`
	Created                 time.Time   `json:"created"`
	PracticeLength          int         `json:"practice_length"`
	QualifyLength           int         `json:"qualify_length"`
	QualifyLaps             int         `json:"qualify_laps"`
	RaceLength              int         `json:"race_length"`
	RaceLaps                int         `json:"race_laps"`
	HeatRace                bool        `json:"heat_race"`
	Host                    CommonHost  `json:"host"`
	Cars                    []struct {
		CarID              int    `json:"car_id"`
		CarName            string `json:"car_name"`
		CarClassID         int    `json:"car_class_id"`
		CarClassName       string `json:"car_class_name"`
		CarClassShortName  string `json:"car_class_short_name"`
		CarNameAbbreviated string `json:"car_name_abbreviated"`
	} `json:"cars"`
}
