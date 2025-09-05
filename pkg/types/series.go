package types

import "time"

type SeriesSeasonsReq struct {
	IncludeSeries bool `url:"include_series,omitempty"`
}

type SeriesSeasonsResp struct {
	SeasonID     int    `json:"season_id"`
	SeasonName   string `json:"season_name"`
	Active       bool   `json:"active"`
	CarClassIds  []int  `json:"car_class_ids"`
	CarSwitching bool   `json:"car_switching"`
	CarTypes     []struct {
		CarType string `json:"car_type"`
	} `json:"car_types"`
	CautionLapsDoNotCount    bool `json:"caution_laps_do_not_count"`
	Complete                 bool `json:"complete"`
	ConnectionBlackFlag      bool `json:"connection_black_flag"`
	ConsecCautionWithinNlaps int  `json:"consec_caution_within_nlaps"`
	ConsecCautionsSingleFile bool `json:"consec_cautions_single_file"`
	CrossLicense             bool `json:"cross_license"`
	DistributedMatchmaking   bool `json:"distributed_matchmaking"`
	DriverChangeRule         int  `json:"driver_change_rule"`
	DriverChanges            bool `json:"driver_changes"`
	Drops                    int  `json:"drops"`
	EnablePitlaneCollisions  bool `json:"enable_pitlane_collisions"`
	FixedSetup               bool `json:"fixed_setup"`
	GreenWhiteCheckeredLimit int  `json:"green_white_checkered_limit"`
	GridByClass              bool `json:"grid_by_class"`
	HardcoreLevel            int  `json:"hardcore_level"`
	HasSupersessions         bool `json:"has_supersessions"`
	IgnoreLicenseForPractice bool `json:"ignore_license_for_practice"`
	IncidentLimit            int  `json:"incident_limit"`
	IncidentWarnMode         int  `json:"incident_warn_mode"`
	IncidentWarnParam1       int  `json:"incident_warn_param1"`
	IncidentWarnParam2       int  `json:"incident_warn_param2"`
	IsHeatRacing             bool `json:"is_heat_racing"`
	LicenseGroup             int  `json:"license_group"`
	LicenseGroupTypes        []struct {
		LicenseGroupType int `json:"license_group_type"`
	} `json:"license_group_types"`
	LuckyDog                   bool   `json:"lucky_dog"`
	MaxTeamDrivers             int    `json:"max_team_drivers"`
	MaxWeeks                   int    `json:"max_weeks"`
	MinTeamDrivers             int    `json:"min_team_drivers"`
	Multiclass                 bool   `json:"multiclass"`
	MustUseDiffTireTypesInRace bool   `json:"must_use_diff_tire_types_in_race"`
	NumFastTows                int    `json:"num_fast_tows"`
	NumOptLaps                 int    `json:"num_opt_laps"`
	Official                   bool   `json:"official"`
	OpDuration                 int    `json:"op_duration"`
	OpenPracticeSessionTypeID  int    `json:"open_practice_session_type_id"`
	QualifierMustStartRace     bool   `json:"qualifier_must_start_race"`
	RaceWeek                   int    `json:"race_week"`
	RaceWeekToMakeDivisions    int    `json:"race_week_to_make_divisions"`
	RegUserCount               int    `json:"reg_user_count"`
	RegionCompetition          bool   `json:"region_competition"`
	RestrictByMember           bool   `json:"restrict_by_member"`
	RestrictToCar              bool   `json:"restrict_to_car"`
	RestrictViewing            bool   `json:"restrict_viewing"`
	ScheduleDescription        string `json:"schedule_description"`
	Schedules                  []struct {
		SeasonID        int `json:"season_id"`
		RaceWeekNum     int `json:"race_week_num"`
		CarRestrictions []struct {
			CarID           int     `json:"car_id"`
			MaxDryTireSets  int     `json:"max_dry_tire_sets"`
			MaxPctFuelFill  int     `json:"max_pct_fuel_fill"`
			PowerAdjustPct  float64 `json:"power_adjust_pct"`
			WeightPenaltyKg float64 `json:"weight_penalty_kg"`
		} `json:"car_restrictions"`
		Category                string `json:"category"`
		CategoryID              int    `json:"category_id"`
		EnablePitlaneCollisions bool   `json:"enable_pitlane_collisions"`
		FullCourseCautions      bool   `json:"full_course_cautions"`
		PracticeLength          int    `json:"practice_length"`
		QualAttached            bool   `json:"qual_attached"`
		QualifyLaps             int    `json:"qualify_laps"`
		QualifyLength           int    `json:"qualify_length"`
		RaceLapLimit            int    `json:"race_lap_limit"`
		RaceTimeDescriptors     []struct {
			DayOffset        []int  `json:"day_offset"`
			FirstSessionTime string `json:"first_session_time"`
			RepeatMinutes    int    `json:"repeat_minutes"`
			Repeating        bool   `json:"repeating"`
			SessionMinutes   int    `json:"session_minutes"`
			StartDate        string `json:"start_date"`
			SuperSession     bool   `json:"super_session"`
		} `json:"race_time_descriptors"`
		// RaceTimeLimit       interface{}   `json:"race_time_limit"`
		// RaceWeekCarClassIds []interface{} `json:"race_week_car_class_ids"`
		// RaceWeekCars        []interface{} `json:"race_week_cars"`
		RestartType    string `json:"restart_type"`
		ScheduleName   string `json:"schedule_name"`
		SeasonName     string `json:"season_name"`
		SeriesID       int    `json:"series_id"`
		SeriesName     string `json:"series_name"`
		ShortParadeLap bool   `json:"short_parade_lap"`
		// SpecialEventType    interface{}   `json:"special_event_type"`
		StartDate string `json:"start_date"`
		StartType string `json:"start_type"`
		StartZone bool   `json:"start_zone"`
		Track     struct {
			Category   string `json:"category"`
			CategoryID int    `json:"category_id"`
			ConfigName string `json:"config_name"`
			TrackID    int    `json:"track_id"`
			TrackName  string `json:"track_name"`
		} `json:"track"`
		TrackState struct {
			LeaveMarbles bool `json:"leave_marbles"`
		} `json:"track_state"`
		WarmupLength int `json:"warmup_length"`
		Weather      struct {
			AllowFog        bool `json:"allow_fog"`
			ForecastOptions struct {
				AllowFog      bool  `json:"allow_fog"`
				ForecastType  int   `json:"forecast_type"`
				Precipitation int   `json:"precipitation"`
				Skies         int   `json:"skies"`
				StopPrecip    int   `json:"stop_precip"`
				Temperature   int   `json:"temperature"`
				WeatherSeed   int64 `json:"weather_seed"`
				WindDir       int   `json:"wind_dir"`
				WindSpeed     int   `json:"wind_speed"`
			} `json:"forecast_options"`
			PrecipOption            int    `json:"precip_option"`
			RelHumidity             int    `json:"rel_humidity"`
			SimulatedStartTime      string `json:"simulated_start_time"`
			SimulatedTimeMultiplier int    `json:"simulated_time_multiplier"`
			SimulatedTimeOffsets    []int  `json:"simulated_time_offsets"`
			Skies                   int    `json:"skies"`
			TempUnits               int    `json:"temp_units"`
			TempValue               int    `json:"temp_value"`
			TimeOfDay               int    `json:"time_of_day"`
			TrackWater              int    `json:"track_water"`
			Version                 int    `json:"version"`
			WeatherSummary          struct {
				MaxPrecipRate     float64 `json:"max_precip_rate"`
				MaxPrecipRateDesc string  `json:"max_precip_rate_desc"`
				PrecipChance      float64 `json:"precip_chance"`
				SkiesHigh         float64 `json:"skies_high"`
				SkiesLow          float64 `json:"skies_low"`
				TempHigh          float64 `json:"temp_high"`
				TempLow           float64 `json:"temp_low"`
				TempUnits         int     `json:"temp_units"`
				WindDir           int     `json:"wind_dir"`
				WindHigh          float64 `json:"wind_high"`
				WindLow           float64 `json:"wind_low"`
				WindUnits         float64 `json:"wind_units"`
			} `json:"weather_summary"`
			WeatherURL string `json:"weather_url"`
			WindDir    int    `json:"wind_dir"`
			WindUnits  int    `json:"wind_units"`
			WindValue  int    `json:"wind_value"`
		} `json:"weather"`
		WeekEndTime time.Time `json:"week_end_time"`
	} `json:"schedules"`
	SeasonQuarter      int       `json:"season_quarter"`
	SeasonShortName    string    `json:"season_short_name"`
	SeasonYear         int       `json:"season_year"`
	SendToOpenPractice bool      `json:"send_to_open_practice"`
	SeriesID           int       `json:"series_id"`
	ShortParadeLap     bool      `json:"short_parade_lap"`
	StartDate          time.Time `json:"start_date"`
	StartOnQualTire    bool      `json:"start_on_qual_tire"`
	StartZone          bool      `json:"start_zone"`
	TrackTypes         []struct {
		TrackType string `json:"track_type"`
	} `json:"track_types"`
	UnsportConductRuleMode int `json:"unsport_conduct_rule_mode"`
	HeatSesInfo            struct {
		ConsolationDeltaMaxFieldSize         int       `json:"consolation_delta_max_field_size"`
		ConsolationDeltaSessionLaps          int       `json:"consolation_delta_session_laps"`
		ConsolationDeltaSessionLengthMinutes int       `json:"consolation_delta_session_length_minutes"`
		ConsolationFirstMaxFieldSize         int       `json:"consolation_first_max_field_size"`
		ConsolationFirstSessionLaps          int       `json:"consolation_first_session_laps"`
		ConsolationFirstSessionLengthMinutes int       `json:"consolation_first_session_length_minutes"`
		ConsolationNumPositionToInvert       int       `json:"consolation_num_position_to_invert"`
		ConsolationNumToConsolation          int       `json:"consolation_num_to_consolation"`
		ConsolationNumToMain                 int       `json:"consolation_num_to_main"`
		ConsolationRunAlways                 bool      `json:"consolation_run_always"`
		ConsolationScoresChampPoints         bool      `json:"consolation_scores_champ_points"`
		Created                              time.Time `json:"created"`
		CustID                               int       `json:"cust_id"`
		Description                          string    `json:"description"`
		HeatCautionType                      int       `json:"heat_caution_type"`
		HeatInfoID                           int       `json:"heat_info_id"`
		HeatInfoName                         string    `json:"heat_info_name"`
		HeatLaps                             int       `json:"heat_laps"`
		HeatLengthMinutes                    int       `json:"heat_length_minutes"`
		HeatMaxFieldSize                     int       `json:"heat_max_field_size"`
		HeatNumFromEachToMain                int       `json:"heat_num_from_each_to_main"`
		HeatNumPositionToInvert              int       `json:"heat_num_position_to_invert"`
		HeatScoresChampPoints                bool      `json:"heat_scores_champ_points"`
		HeatSessionMinutesEstimate           int       `json:"heat_session_minutes_estimate"`
		Hidden                               bool      `json:"hidden"`
		MainLaps                             int       `json:"main_laps"`
		MainLengthMinutes                    int       `json:"main_length_minutes"`
		MainMaxFieldSize                     int       `json:"main_max_field_size"`
		MainNumPositionToInvert              int       `json:"main_num_position_to_invert"`
		MaxEntrants                          int       `json:"max_entrants"`
		OpenPractice                         bool      `json:"open_practice"`
		PreMainPracticeLengthMinutes         int       `json:"pre_main_practice_length_minutes"`
		PreQualNumToMain                     int       `json:"pre_qual_num_to_main"`
		PreQualPracticeLengthMinutes         int       `json:"pre_qual_practice_length_minutes"`
		QualCautionType                      int       `json:"qual_caution_type"`
		QualLaps                             int       `json:"qual_laps"`
		QualLengthMinutes                    int       `json:"qual_length_minutes"`
		QualNumToMain                        int       `json:"qual_num_to_main"`
		QualOpenDelaySeconds                 int       `json:"qual_open_delay_seconds"`
		QualScoresChampPoints                bool      `json:"qual_scores_champ_points"`
		QualScoring                          int       `json:"qual_scoring"`
		QualStyle                            int       `json:"qual_style"`
		RaceStyle                            int       `json:"race_style"`
	} `json:"heat_ses_info,omitempty"`
	RegOpenMinutes int `json:"reg_open_minutes,omitempty"`
}


// StatsSeries

type SeriesStatsSeriesReq struct {
	Official bool `url:"official,omitempty"`
}

type SeriesStatsSeriesResp struct {
	SeriesID          int         `json:"series_id"`
	SeriesName        string      `json:"series_name"`
	SeriesShortName   string      `json:"series_short_name"`
	CategoryID        int         `json:"category_id"`
	Category          string      `json:"category"`
	Active            bool        `json:"active"`
	Official          bool        `json:"official"`
	FixedSetup        bool        `json:"fixed_setup"`
	Logo              interface{} `json:"logo"`
	LicenseGroup      int         `json:"license_group"`
	LicenseGroupTypes []struct {
		LicenseGroupType int `json:"license_group_type"`
	} `json:"license_group_types"`
	AllowedLicenses []struct {
		GroupName       string `json:"group_name"`
		LicenseGroup    int    `json:"license_group"`
		MaxLicenseLevel int    `json:"max_license_level"`
		MinLicenseLevel int    `json:"min_license_level"`
		ParentID        int    `json:"parent_id"`
	} `json:"allowed_licenses"`
	Seasons []struct {
		SeasonID          int    `json:"season_id"`
		SeriesID          int    `json:"series_id"`
		SeasonName        string `json:"season_name"`
		SeasonShortName   string `json:"season_short_name"`
		SeasonYear        int    `json:"season_year"`
		SeasonQuarter     int    `json:"season_quarter"`
		Active            bool   `json:"active"`
		Official          bool   `json:"official"`
		DriverChanges     bool   `json:"driver_changes"`
		FixedSetup        bool   `json:"fixed_setup"`
		LicenseGroup      int    `json:"license_group"`
		HasSupersessions  bool   `json:"has_supersessions"`
		CarSwitching      bool   `json:"car_switching"`
		LicenseGroupTypes []struct {
			LicenseGroupType int `json:"license_group_type"`
		} `json:"license_group_types"`
		CarClasses []struct {
			CarClassID    int    `json:"car_class_id"`
			ShortName     string `json:"short_name"`
			Name          string `json:"name"`
			RelativeSpeed int    `json:"relative_speed"`
		} `json:"car_classes"`
		RaceWeeks []struct {
			SeasonID    int `json:"season_id"`
			RaceWeekNum int `json:"race_week_num"`
			Track       struct {
				ConfigName string `json:"config_name"`
				TrackID    int    `json:"track_id"`
				TrackName  string `json:"track_name"`
			} `json:"track"`
		} `json:"race_weeks"`
	} `json:"seasons"`
	SearchFilters string `json:"search_filters,omitempty"`
}