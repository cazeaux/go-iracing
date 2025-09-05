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

// EventLog -- TODO

/*
{
  "success": true,
  "session_info": {
    "subsession_id": 70845764,
    "session_id": 247522669,
    "simsession_number": 0,
    "simsession_type": 6,
    "simsession_name": "RACE",
    "event_type": 5,
    "event_type_name": "Race",
    "private_session_id": -1,
    "season_name": "IMSA iRacing Series - 2024 Season 3",
    "season_short_name": "2024 Season 3",
    "series_name": "IMSA iRacing Series",
    "series_short_name": "IMSA iRacing Series",
    "start_time": "2024-08-18T08:45:00Z",
    "track": {
      "config_name": "24 Heures du Mans",
      "track_id": 268,
      "track_name": "Circuit des 24 Heures du Mans"
    }
  },
  "chunk_info": {
    "chunk_size": 500,
    "num_chunks": 1,
    "rows": 13,
    "base_download_url": "https://scorpio-assets.s3.amazonaws.com/members/messaging-services/long_lived/lapdata/subsession/70845764/0/eventlog/",
    "chunk_file_names": [
      "ece964167510f8c207f4f023eec9fa60981701335710bea88d5251983e67aa16.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T063958Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=3aed180d21dbf761c0e0cfb975dc43a6d15fd95fa08b9ff372cbb951944bc8be"
    ]
  }
}
*/

// Get

type ResultsGetReq struct {
	SubsessionID    int  `url:"subsession_id"`
	IncludeLicenses bool `url:"include_licenses,omitempty"`
}

type ResultsGetResp struct {
	SubsessionID    int `json:"subsession_id"`
	AllowedLicenses []struct {
		GroupName       string `json:"group_name"`
		LicenseGroup    int    `json:"license_group"`
		MaxLicenseLevel int    `json:"max_license_level"`
		MinLicenseLevel int    `json:"min_license_level"`
		ParentID        int    `json:"parent_id"`
	} `json:"allowed_licenses"`
	AssociatedSubsessionIds []int `json:"associated_subsession_ids"`
	CanProtest              bool  `json:"can_protest"`
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
		SimsessionNumber   int    `json:"simsession_number"`
		SimsessionName     string `json:"simsession_name"`
		SimsessionType     int    `json:"simsession_type"`
		SimsessionTypeName string `json:"simsession_type_name"`
		SimsessionSubtype  int    `json:"simsession_subtype"`
		WeatherResult      struct {
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
		} `json:"weather_result"`
		Results []struct {
			CustID                int       `json:"cust_id"`
			DisplayName           string    `json:"display_name"`
			AggregateChampPoints  int       `json:"aggregate_champ_points"`
			Ai                    bool      `json:"ai"`
			AverageLap            int       `json:"average_lap"`
			BestLapNum            int       `json:"best_lap_num"`
			BestLapTime           int       `json:"best_lap_time"`
			BestNlapsNum          int       `json:"best_nlaps_num"`
			BestNlapsTime         int       `json:"best_nlaps_time"`
			BestQualLapAt         time.Time `json:"best_qual_lap_at"`
			BestQualLapNum        int       `json:"best_qual_lap_num"`
			BestQualLapTime       int       `json:"best_qual_lap_time"`
			CarClassID            int       `json:"car_class_id"`
			CarClassName          string    `json:"car_class_name"`
			CarClassShortName     string    `json:"car_class_short_name"`
			CarID                 int       `json:"car_id"`
			CarName               string    `json:"car_name"`
			Carcfg                int       `json:"carcfg"`
			ChampPoints           int       `json:"champ_points"`
			ClassInterval         int       `json:"class_interval"`
			CountryCode           string    `json:"country_code"`
			Division              int       `json:"division"`
			DivisionName          string    `json:"division_name"`
			DropRace              bool      `json:"drop_race"`
			FinishPosition        int       `json:"finish_position"`
			FinishPositionInClass int       `json:"finish_position_in_class"`
			FlairID               int       `json:"flair_id"`
			FlairName             string    `json:"flair_name"`
			FlairShortname        string    `json:"flair_shortname"`
			Friend                bool      `json:"friend"`
			Helmet                struct {
				Pattern    int    `json:"pattern"`
				Color1     string `json:"color1"`
				Color2     string `json:"color2"`
				Color3     string `json:"color3"`
				FaceType   int    `json:"face_type"`
				HelmetType int    `json:"helmet_type"`
			} `json:"helmet"`
			Incidents         int `json:"incidents"`
			Interval          int `json:"interval"`
			LapsComplete      int `json:"laps_complete"`
			LapsLead          int `json:"laps_lead"`
			LeagueAggPoints   int `json:"league_agg_points"`
			LeaguePoints      int `json:"league_points"`
			LicenseChangeOval int `json:"license_change_oval"`
			LicenseChangeRoad int `json:"license_change_road"`
			Livery            struct {
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
			MaxPctFuelFill          int     `json:"max_pct_fuel_fill"`
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
		} `json:"results"`
	} `json:"session_results"`
	SessionSplits []struct {
		SubsessionID         int `json:"subsession_id"`
		EventStrengthOfField int `json:"event_strength_of_field"`
	} `json:"session_splits"`
	SpecialEventType int       `json:"special_event_type"`
	StartTime        time.Time `json:"start_time"`
	Track            struct {
		Category   string `json:"category"`
		CategoryID int    `json:"category_id"`
		ConfigName string `json:"config_name"`
		TrackID    int    `json:"track_id"`
		TrackName  string `json:"track_name"`
	} `json:"track"`
	TrackState struct {
		LeaveMarbles   bool `json:"leave_marbles"`
		PracticeRubber int  `json:"practice_rubber"`
		QualifyRubber  int  `json:"qualify_rubber"`
		RaceRubber     int  `json:"race_rubber"`
		WarmupRubber   int  `json:"warmup_rubber"`
	} `json:"track_state"`
	Weather struct {
		AllowFog                      bool   `json:"allow_fog"`
		Fog                           int    `json:"fog"`
		PrecipMm2HrBeforeFinalSession int    `json:"precip_mm2hr_before_final_session"`
		PrecipMmFinalSession          int    `json:"precip_mm_final_session"`
		PrecipOption                  int    `json:"precip_option"`
		PrecipTimePct                 int    `json:"precip_time_pct"`
		RelHumidity                   int    `json:"rel_humidity"`
		SimulatedStartTime            string `json:"simulated_start_time"`
		Skies                         int    `json:"skies"`
		TempUnits                     int    `json:"temp_units"`
		TempValue                     int    `json:"temp_value"`
		TimeOfDay                     int    `json:"time_of_day"`
		TrackWater                    int    `json:"track_water"`
		Type                          int    `json:"type"`
		Version                       int    `json:"version"`
		WeatherVarInitial             int    `json:"weather_var_initial"`
		WeatherVarOngoing             int    `json:"weather_var_ongoing"`
		WindDir                       int    `json:"wind_dir"`
		WindUnits                     int    `json:"wind_units"`
		WindValue                     int    `json:"wind_value"`
	} `json:"weather"`
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

// lap_chart_data -- TODO

/*
{
  "success": true,
  "session_info": {
    "subsession_id": 70845764,
    "session_id": 247522669,
    "simsession_number": 0,
    "simsession_type": 6,
    "simsession_name": "RACE",
    "num_laps_for_qual_average": 2,
    "num_laps_for_solo_average": 2,
    "event_type": 5,
    "event_type_name": "Race",
    "private_session_id": -1,
    "season_name": "IMSA iRacing Series - 2024 Season 3",
    "season_short_name": "2024 Season 3",
    "series_name": "IMSA iRacing Series",
    "series_short_name": "IMSA iRacing Series",
    "start_time": "2024-08-18T08:45:00Z",
    "track": {
      "config_name": "24 Heures du Mans",
      "track_id": 268,
      "track_name": "Circuit des 24 Heures du Mans"
    }
  },
  "best_lap_num": 4,
  "best_lap_time": 2035896,
  "best_nlaps_num": -1,
  "best_nlaps_time": -1,
  "best_qual_lap_num": -1,
  "best_qual_lap_time": -1,
  "best_qual_lap_at": null,
  "chunk_info": {
    "chunk_size": 500,
    "num_chunks": 1,
    "rows": 335,
    "base_download_url": "https://scorpio-assets.s3.amazonaws.com/members/messaging-services/long_lived/lapdata/subsession/70845764/0/chart/",
    "chunk_file_names": [
      "884dcf61ed62a74ee01b6429486a5c0b872aec1d66a76df30099885fa35536ad.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070618Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=ce1fa1b7d0ee08117e03a11d1623fb76e5f4ff30ad3ed30dd9095eddb03c74c9"
    ]
  },
  "last_updated": "2025-09-13T07:06:18.757579751Z"
}



{
  "success": true,
  "session_info": {
    "subsession_id": 78354169,
    "session_id": 279907628,
    "simsession_number": 0,
    "simsession_type": 6,
    "simsession_name": "RACE",
    "num_laps_for_qual_average": 2,
    "num_laps_for_solo_average": 4,
    "event_type": 5,
    "event_type_name": "Race",
    "private_session_id": -1,
    "season_name": "2025 24 hours of Spa Presented by Falken Tyre",
    "season_short_name": "2025 Spa 24",
    "series_name": "24 Hours of Spa",
    "series_short_name": "24 Hours of Spa",
    "start_time": "2025-07-12T12:00:00Z",
    "track": {
      "config_name": "Endurance",
      "track_id": 525,
      "track_name": "Circuit de Spa-Francorchamps"
    }
  },
  "best_lap_num": 310,
  "best_lap_time": 1354196,
  "best_nlaps_num": -1,
  "best_nlaps_time": -1,
  "best_qual_lap_num": -1,
  "best_qual_lap_time": -1,
  "best_qual_lap_at": null,
  "chunk_info": {
    "chunk_size": 500,
    "num_chunks": 43,
    "rows": 21206,
    "base_download_url": "https://scorpio-assets.s3.amazonaws.com/members/messaging-services/long_lived/lapdata/subsession/78354169/0/chart/",
    "chunk_file_names": [
      "050b4b424143c63b8f3e79b0292619e850ee03c567737080518e53983e3a463e.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=98eadba80d4bb24623c89ef02cb1797f3cdfe8306c646103bd32fc42101b5711",
      "c5be737c38d487ab1d5243d1e4b03d8d3a565dafb245be3de95db5695d69e153.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=c89ca692f84396348aaa011cd9a7af5e6881c6de4228313c7fed8d96398982e6",
      "fa96e2c04ab75d67843e9bb74deb7324f8138d2dfd119133429696e01b3a1d4d.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=0d14b92c7baf3fe91edca5f5f2c4da2ca16da90dd77c8bc5a89ecf8f1079fd43",
      "ae376394b2492d1b5e679be9df8ae9e701ec6a69567944903d4090dadde47062.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=88486a7650a1610a02d8384b2fd9ff3098a8e9ef4fceeaaffbd8490b0e4d819c",
      "2218329a76c9a2edf1406f9010b9e4ab2b7d31828b31768893061edc75dc98b4.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=9f93f58502bc65b801ae04128f1b4eb0fce484f61f64ad941838865b8a548178",
      "094868d185826a1e0b7a57682df11f249ff5d31cc6dc7ee6c72529d681d3ca43.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=ae356579de790c8934dc6f285f5e17a4d157b2815f9d74648b59a416a7e3d253",
      "63cf7673fcbd309d43f946fb88269b512348794bed7315faa5702a95a0ad8456.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=2fdc6a0949596765f641fc90eaf117ef0f072b5cd4b1ff4924ec792728f6b034",
      "9109ff313d1fbe273cbd9e4effca648a510a0ebb6d929f1c723785ccbc066645.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=021cf8a363b51507ad6e601f734180e0c8b2d478e8257c6833a280afed847a97",
      "2ba6cba210ba8de06dee6b0660a8e930130ed1f8857de4b97a162bc77b6a48bd.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=c88004699a80373aa4abcfbf9ef9d8cf79aaa7b0d9e6d3104905a2969e10a6f9",
      "3622f43724f44fc9b5f03017626d3ac881ba86d1ec0ba2f8014cba40bd482fb6.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=232efe8c28ba9775f9bb99ac578034e973b40e215bc5f7c2c40ebff9a5264ba2",
      "5134750e3494d2d22a18d961e736aabcf8044fa60b6edcab909de1c8ac91f294.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=f046f260c7454c8e824c905ed466a8b4db2fdee9af69a754541ace1c8d8a7593",
      "b90fd9cafb24f17384132bb546400dd1855d73225cde5403b2d6d2c30f678ce3.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=e03004103f5bb0b0d4e994135f4ad832b763ebbbca5fde8846667c916cc9ea1b",
      "b3deaaf4b866e3ab5907185460378900adbee44c5818af524c06848147fdbe38.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=a9d16621caf7b281bcf23709ac0b08a0ce5c089a9659711a556d0c06ca7d9af3",
      "ab2d8c032790113336103c08ca8f72ad4b3714e63a671ee74fbfb94c05df2146.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=adf81757f5f1bda8ae6d0887d0715dbfaf584147f2981afa349f9113bc37c703",
      "541a0a713ea8e7e41118b86187edef03e8fad99bbadcc27f336518d38682ce18.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=3f5ce5d46380d2d1e13140064edacfbf853f685529693ee3ef60543fc6ba7fda",
      "5eb9dd7a04d0c0876ed9dd491cd56fe9b7cb8335e4f68f03b74915132a887d37.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=41d121be947d7df135a60b146508677716470cb979de19eee30fd91216f8b1d1",
      "a2ff9af3372f1736163be1ccd3c2dd3171407ffed6f0183a2056ea9430069cde.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=fd08e51cc5ea05fcb2b9ca27d0d7b249dbe01eecfd2330f47d8b7ec3b8ad2b14",
      "356ce9ffc325068e6609bc8a5a9163982ad99d58682753e7c90a074654379a27.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=0d2cb198b688f1a7644c4c11b952734ed255c6694ce44c2a63429d7159c504c2",
      "db237df185abe6782123b40777de0d3161c6892d2a13333bc23f500e6d19b250.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=205512574a226aba6625927986cc2c3b286180532239b2f51c22b6aed39dd198",
      "16b837a5ead79165c36cffe846bf8a398dc247ecd531f9a923dddfe75b6045d3.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=1cf9dd8373706b9bdbd87a24cbe4d1280029977d7afc1f767f58273b611ee2a4",
      "8cb524fb25cc8e21c8d33bd7a6aede09fd3a652fae25772f2bc1ecf1d60f727a.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=89e86d7fe37a1288a5ee458e02a1b47ae21ccb539104aae1179f9a46baa39a10",
      "bd9281fb0e1965b8b7f90b8d6b2d69886e35bf67859240127e3002c353bb8dc2.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=ef4898860cdcde9ca31c6e7ecd5535b51065afd98017ef0f2cff228e6018877c",
      "dfd572485dd23f56b7e8928e86ef8ce2fd1d33c41ec001ef80b19f1ab0843869.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=ca846210dfe12789e96055fbfd80e89035fbb30effe973f74647f7e934b27677",
      "813854c517fd9fa943f9fbdce54c1f6447915c2f3f7c9ac1b4e39866e391c5ce.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=f5b64c6328ce273905c3841105b2d91ac71f4a71d28f2c517a64ed7134052dc1",
      "94d023c043c1c536855cfcbce42feb73d4de13257bb5446b9ff7df772b5fefff.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=4234062f379c1544a7300729cfe8523cb0cc26c8de7ea067c4b80b9c11914256",
      "b74217dadaf081843e1e04dc88badd5a93cddd11572d9bb1de6202c5e5cefa9a.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=933039e633d5e11303a6d0dfae9aa3299824ad53426e69bbabfade43124a2f62",
      "eed75c796d42a3f35aa052620c63d415a777bfa5664ae0c774e1401a25d8cdb7.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=bcea1713361e303cd37aef49b6d22f0ca85e0d8a1f237708a531054fbcf26337",
      "cf0bf0181395f17fa202b1227e096deea15fc82f48ef2173f8b30e8e7b87c74d.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=e6511f72caaaf56b4060e1b23be2e4ecef33adee1313b3a36090dde61f86fe0e",
      "06cf4c16c9477c51ea235b8806a000fe4bf35027df20ee34a0c3625b91d57698.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=6a6575c1e9b0a50d7dd6a85dcb8a748e4c7d02b71c89f816174b45a1677e2f78",
      "a125816331dab8c145bdb40e1dfb16a957b2a7d49f47e04a027361c8e562c249.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=fca0ac99dac094f93d219ae46a770446fc5ba4bd0a0ab2fb0308eae38aa2db37",
      "a69459f520bd8ea0dcb67e82ed18628c220a4c1046cf47cd3d9f1422af0a0a65.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=b787c2e7745ce64926ac2870b4baa13b4b9a16785fa8c8e2545014e1ce96b24f",
      "2e09032f3ed27cae65a89a44539981c642815e26f678c081204f44c681ff3bec.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=b998ff5717cce9731454b09ad50a87f7f27b41fd404821b995c5616903b49f51",
      "6b47f260a1bc9099a83a09a655fac584abd43809d4682b4f251d8c93f02a993e.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=6272918a15196d2b6f0fa556556a85e3a22a9562c62874e67bc494e2ac4df349",
      "68ec43d3be7f440ab8a0c29326557d30bd3aefc75e8dc44056d576224aded208.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=703e85b69bd9785a2377602398a2d36d0fab2208420452febd9848909b5295e9",
      "fecbbda0307898c3d5ff61771b14601db2f0b7f6a88ee84b4d0b3c8ad11893a3.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=04990707494a2199f7cb0c5b22c651f3ef9e172531e0ea5b3118c5278b182526",
      "e47b9e7d9f8c31084baddde182058b7250fc10a6fb81405587ef930f0fc79d5a.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=ac7bf21480169f8f095e9acc70c264ae64fb1f14f4edfb874099e90db5dd5652",
      "d613dc02ebe3a669068f6cdc3fe70f50370f56f11ba68f8c6ee826771e68673f.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=fe02b74a00f1180d6ab6287c37acb24e3af9ea68f2789b22730967107666c5c6",
      "ea2f1e60a034a7ae67670a46f147c9c3aa5d0874013a95f0418efb80c8c7d43e.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=8e4e4dec45f08c85aa02872e690def2bf54cc7caceec0284884a6b09498dead5",
      "c052ee43fe6895fa580d0948c42012e08d93f4a5afd4f08aed1465de4ab46bdb.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=dcc1e4e027e1d960bb8185f6e50c79b064bbeaf29abe00513aa19f03da52a6e9",
      "01257f00a2f176b761b1cad066d29dc2dfebaebdad562dab96be7cfe6a2f4997.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=d8001ca5925c961654be5b695ac00b862a742eb6b41661f00b8ea6a90663c264",
      "bcf5db97039e9ea9b2f22126a63ee78bc1237d5eb1c11662f40ee273c78d4fb1.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=205a896e6075f61a5e7128cd30e29dbdefc32f80d88db76f607384351c9479d3",
      "55522e4acf51d1a3f015dd7328f7922f6a9f09b3db9202bde0e730c9e1959f7b.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=9e19b9d0a3d335412a118d940cdd0a0aaa1cbf4781c33c3eefdb446d58431480",
      "e0a2f5269e6999cc007eec744e99d1689755bdc27ab474660f7040dffbde07d6.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250913T070752Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250913%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=e6c336a1fd016b3c89ef86626b4cba303fc3e5c89764a97768791fec8d7d2467"
    ]
  },
  "last_updated": "2025-09-13T07:07:52.668467488Z"
}
*/