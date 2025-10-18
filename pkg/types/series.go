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
		Category                string                     `json:"category"`
		CategoryID              int                        `json:"category_id"`
		EnablePitlaneCollisions bool                       `json:"enable_pitlane_collisions"`
		FullCourseCautions      bool                       `json:"full_course_cautions"`
		PracticeLength          int                        `json:"practice_length"`
		QualAttached            bool                       `json:"qual_attached"`
		QualifyLaps             int                        `json:"qualify_laps"`
		QualifyLength           int                        `json:"qualify_length"`
		RaceLapLimit            int                        `json:"race_lap_limit"`
		RaceTimeDescriptors     []CommonSessTimeDescriptor `json:"race_time_descriptors"`
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
		WarmupLength int               `json:"warmup_length"`
		Weather      CommonWeatherFull `json:"weather"`
		WeekEndTime  time.Time         `json:"week_end_time"`
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
	UnsportConductRuleMode int               `json:"unsport_conduct_rule_mode"`
	HeatSesInfo            CommonHeatSesInfo `json:"heat_ses_info,omitempty"`
	RegOpenMinutes         int               `json:"reg_open_minutes,omitempty"`
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
	AllowedLicenses []CommonAllowedLicenses `json:"allowed_licenses"`
	Seasons         []struct {
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

// Assets

type SeriesAssetsReq struct{}

type SeriesAssetsResp map[string]SeriesAsset

type SeriesAsset struct {
	LargeImage string `json:"large_image"`
	Logo       string `json:"logo"`
	SeriesCopy string `json:"series_copy"`
	SeriesID   int    `json:"series_id"`
	SmallImage string `json:"small_image"`
}

// Get

type SeriesGetReq struct{}

type SeriesGetResp struct {
	AllowedLicenses []CommonAllowedLicenses `json:"allowed_licenses"`
	Category        string                  `json:"category"`
	CategoryID      int                     `json:"category_id"`
	Eligible        bool                    `json:"eligible"`
	FirstSeason     struct {
		SeasonYear    int `json:"season_year"`
		SeasonQuarter int `json:"season_quarter"`
	} `json:"first_season"`
	ForumURL        string `json:"forum_url,omitempty"`
	MaxStarters     int    `json:"max_starters"`
	MinStarters     int    `json:"min_starters"`
	OvalCautionType int    `json:"oval_caution_type"`
	RoadCautionType int    `json:"road_caution_type"`
	SeriesID        int    `json:"series_id"`
	SeriesName      string `json:"series_name"`
	SeriesShortName string `json:"series_short_name"`
	SearchFilters   string `json:"search_filters,omitempty"`
}

// Past Seasons

type SeriesPastSeasonsReq struct {
	SeriesID int `url:"series_id"`
}
type SeriesPastSeasonsResp struct {
	Success bool `json:"success"`
	Series  struct {
		SeriesID          int    `json:"series_id"`
		SeriesName        string `json:"series_name"`
		SeriesShortName   string `json:"series_short_name"`
		CategoryID        int    `json:"category_id"`
		Category          string `json:"category"`
		Active            bool   `json:"active"`
		Official          bool   `json:"official"`
		FixedSetup        bool   `json:"fixed_setup"`
		SearchFilters     string `json:"search_filters"`
		Logo              string `json:"logo"`
		LicenseGroup      int    `json:"license_group"`
		LicenseGroupTypes []struct {
			LicenseGroupType int `json:"license_group_type"`
		} `json:"license_group_types"`
		AllowedLicenses []CommonAllowedLicenses `json:"allowed_licenses"`
		Seasons         []struct {
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
	} `json:"series"`
	SeriesID int `json:"series_id"`
}

// Season List

type SeriesSeasonListReq struct {
	IncludeSeries bool `url:"include_series,omitempty"`
	SeasonYear    int  `url:"season_year,omitempty"`
	SeasonQuarter int  `url:"season_quarter,omitempty"`
}

type SeriesSeasonListResp struct {
	Seasons []struct {
		SeasonID             int         `json:"season_id"`
		SeasonName           string      `json:"season_name"`
		Active               bool        `json:"active"`
		AllowedSeasonMembers interface{} `json:"allowed_season_members"`
		CarClassIds          []int       `json:"car_class_ids"`
		CarSwitching         bool        `json:"car_switching"`
		CarTypes             []struct {
			CarType string `json:"car_type"`
		} `json:"car_types"`
		CautionLapsDoNotCount    bool `json:"caution_laps_do_not_count"`
		Complete                 bool `json:"complete"`
		ConnectionBlackFlag      bool `json:"connection_black_flag"`
		ConsecCautionWithinNlaps int  `json:"consec_caution_within_nlaps"`
		ConsecCautionsSingleFile bool `json:"consec_cautions_single_file"`
		CrossLicense             bool `json:"cross_license"`
		CurrentWeekSched         struct {
			RaceWeekNum     int         `json:"race_week_num"`
			Track           CommonTrack `json:"track"`
			CarRestrictions []struct {
				CarID           int     `json:"car_id"`
				MaxDryTireSets  int     `json:"max_dry_tire_sets"`
				MaxPctFuelFill  float64 `json:"max_pct_fuel_fill"`
				PowerAdjustPct  float64 `json:"power_adjust_pct"`
				WeightPenaltyKg float64 `json:"weight_penalty_kg"`
			} `json:"car_restrictions"`
			RaceLapLimit  int         `json:"race_lap_limit"`
			RaceTimeLimit interface{} `json:"race_time_limit"`
			PrecipChance  int         `json:"precip_chance"`
			StartType     string      `json:"start_type"`
			CategoryID    int         `json:"category_id"`
		} `json:"current_week_sched"`
		DistributedMatchmaking bool `json:"distributed_matchmaking"`
		DriverChangeRule       int  `json:"driver_change_rule"`
		DriverChanges          bool `json:"driver_changes"`
		Drops                  int  `json:"drops"`
		Elig                   struct {
			OwnCar   bool `json:"own_car"`
			OwnTrack bool `json:"own_track"`
		} `json:"elig"`
		EnablePitlaneCollisions  bool `json:"enable_pitlane_collisions"`
		FixedSetup               bool `json:"fixed_setup"`
		GreenWhiteCheckeredLimit int  `json:"green_white_checkered_limit"`
		GridByClass              bool `json:"grid_by_class"`
		HardcoreLevel            int  `json:"hardcore_level"`
		HasMpr                   bool `json:"has_mpr"`
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
		LuckyDog                   bool      `json:"lucky_dog"`
		MaxTeamDrivers             int       `json:"max_team_drivers"`
		MaxWeeks                   int       `json:"max_weeks"`
		MinTeamDrivers             int       `json:"min_team_drivers"`
		Multiclass                 bool      `json:"multiclass"`
		MustUseDiffTireTypesInRace bool      `json:"must_use_diff_tire_types_in_race"`
		NumFastTows                int       `json:"num_fast_tows"`
		NumOptLaps                 int       `json:"num_opt_laps"`
		Official                   bool      `json:"official"`
		OpDuration                 int       `json:"op_duration"`
		OpenPracticeSessionTypeID  int       `json:"open_practice_session_type_id"`
		QualifierMustStartRace     bool      `json:"qualifier_must_start_race"`
		RaceWeek                   int       `json:"race_week"`
		RaceWeekToMakeDivisions    int       `json:"race_week_to_make_divisions"`
		RegUserCount               int       `json:"reg_user_count"`
		RegionCompetition          bool      `json:"region_competition"`
		RestrictByMember           bool      `json:"restrict_by_member"`
		RestrictToCar              bool      `json:"restrict_to_car"`
		RestrictViewing            bool      `json:"restrict_viewing"`
		ScheduleDescription        string    `json:"schedule_description"`
		SeasonQuarter              int       `json:"season_quarter"`
		SeasonShortName            string    `json:"season_short_name"`
		SeasonYear                 int       `json:"season_year"`
		SendToOpenPractice         bool      `json:"send_to_open_practice"`
		SeriesID                   int       `json:"series_id"`
		ShortParadeLap             bool      `json:"short_parade_lap"`
		StartDate                  time.Time `json:"start_date"`
		StartOnQualTire            bool      `json:"start_on_qual_tire"`
		StartZone                  bool      `json:"start_zone"`
		TrackTypes                 []struct {
			TrackType string `json:"track_type"`
		} `json:"track_types"`
		UnsportConductRuleMode int               `json:"unsport_conduct_rule_mode"`
		HeatSesInfo            CommonHeatSesInfo `json:"heat_ses_info,omitempty"`
		RegOpenMinutes         int               `json:"reg_open_minutes,omitempty"`
	} `json:"seasons"`
}

// Season Schedule

type SeriesSeasonScheduleReq struct {
	SeasonID int `url:"season_id,omitempty"`
}
type SeriesSeasonScheduleResp struct {
	Success   bool `json:"success"`
	SeasonID  int  `json:"season_id"`
	Schedules []struct {
		SeasonID                int                        `json:"season_id"`
		RaceWeekNum             int                        `json:"race_week_num"`
		CarRestrictions         []interface{}              `json:"car_restrictions"`
		Category                string                     `json:"category"`
		CategoryID              int                        `json:"category_id"`
		EnablePitlaneCollisions bool                       `json:"enable_pitlane_collisions"`
		FullCourseCautions      bool                       `json:"full_course_cautions"`
		PracticeLength          int                        `json:"practice_length"`
		QualAttached            bool                       `json:"qual_attached"`
		QualTimeDescriptors     []CommonSessTimeDescriptor `json:"qual_time_descriptors"`
		QualifyLaps             int                        `json:"qualify_laps"`
		QualifyLength           int                        `json:"qualify_length"`
		RaceLapLimit            int                        `json:"race_lap_limit"`
		RaceTimeDescriptors     []CommonSessTimeDescriptor `json:"race_time_descriptors"`
		RaceTimeLimit           int                        `json:"race_time_limit"`
		RaceWeekCarClassIds     []int                      `json:"race_week_car_class_ids"`
		RaceWeekCars            []string                   `json:"race_week_cars"`
		RestartType             string                     `json:"restart_type"`
		ScheduleName            string                     `json:"schedule_name"`
		SeasonName              string                     `json:"season_name"`
		SeriesID                int                        `json:"series_id"`
		SeriesName              string                     `json:"series_name"`
		ShortParadeLap          bool                       `json:"short_parade_lap"`
		SpecialEventType        string                     `json:"special_event_type"`
		StartDate               string                     `json:"start_date"`
		StartType               string                     `json:"start_type"`
		StartZone               bool                       `json:"start_zone"`
		Track                   CommonTrack                `json:"track"`
		TrackState              CommonTrackState           `json:"track_state"`
		WarmupLength            int                        `json:"warmup_length"`
		Weather                 CommonWeatherFull          `json:"weather"`
		WeekEndTime             time.Time                  `json:"week_end_time"`
	} `json:"schedules"`
}
