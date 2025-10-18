package types

import "time"

type HostedCombinedSessionsReq struct {
	PackageID int `url:"package_id"`
}

type HostedSessionsReq struct{}

type HostedCombinedSessionsResp struct {
	Success    bool `json:"success"`
	Subscribed bool `json:"subscribed"`
	Sequence   int  `json:"sequence"`
	Sessions   []struct {
		NumDrivers                        int              `json:"num_drivers"`
		NumSpotters                       int              `json:"num_spotters"`
		NumSpectators                     int              `json:"num_spectators"`
		NumBroadcasters                   int              `json:"num_broadcasters"`
		AvailableReservedBroadcasterSlots int              `json:"available_reserved_broadcaster_slots"`
		NumSpectatorSlots                 int              `json:"num_spectator_slots"`
		AvailableSpectatorSlots           int              `json:"available_spectator_slots"`
		CanBroadcast                      bool             `json:"can_broadcast"`
		CanWatch                          bool             `json:"can_watch"`
		CanSpot                           bool             `json:"can_spot"`
		Elig                              CommonElig       `json:"elig"`
		DriverChanges                     bool             `json:"driver_changes"`
		RestrictViewing                   bool             `json:"restrict_viewing"`
		MaxUsers                          int              `json:"max_users"`
		PrivateSessionID                  int              `json:"private_session_id"`
		SessionID                         int              `json:"session_id"`
		SubsessionID                      int              `json:"subsession_id"`
		PasswordProtected                 bool             `json:"password_protected"`
		SessionName                       string           `json:"session_name"`
		OpenRegExpires                    time.Time        `json:"open_reg_expires"`
		LaunchAt                          time.Time        `json:"launch_at"`
		FullCourseCautions                bool             `json:"full_course_cautions"`
		NumFastTows                       int              `json:"num_fast_tows"`
		RollingStarts                     bool             `json:"rolling_starts"`
		Restarts                          int              `json:"restarts"`
		MulticlassType                    int              `json:"multiclass_type"`
		PitsInUse                         int              `json:"pits_in_use"`
		CarsLeft                          int              `json:"cars_left"`
		MaxDrivers                        int              `json:"max_drivers"`
		HardcoreLevel                     int              `json:"hardcore_level"`
		PracticeLength                    int              `json:"practice_length"`
		LoneQualify                       bool             `json:"lone_qualify"`
		QualifyLaps                       int              `json:"qualify_laps"`
		QualifyLength                     int              `json:"qualify_length"`
		WarmupLength                      int              `json:"warmup_length"`
		RaceLaps                          int              `json:"race_laps"`
		RaceLength                        int              `json:"race_length"`
		TimeLimit                         int              `json:"time_limit"`
		RestrictResults                   bool             `json:"restrict_results"`
		IncidentLimit                     int              `json:"incident_limit"`
		IncidentWarnMode                  int              `json:"incident_warn_mode"`
		IncidentWarnParam1                int              `json:"incident_warn_param1"`
		IncidentWarnParam2                int              `json:"incident_warn_param2"`
		UnsportConductRuleMode            int              `json:"unsport_conduct_rule_mode"`
		ConnectionBlackFlag               bool             `json:"connection_black_flag"`
		LuckyDog                          bool             `json:"lucky_dog"`
		MinTeamDrivers                    int              `json:"min_team_drivers"`
		MaxTeamDrivers                    int              `json:"max_team_drivers"`
		QualifierMustStartRace            bool             `json:"qualifier_must_start_race"`
		DriverChangeRule                  int              `json:"driver_change_rule"`
		FixedSetup                        bool             `json:"fixed_setup"`
		EntryCount                        int              `json:"entry_count"`
		LeagueID                          int              `json:"league_id"`
		LeagueSeasonID                    int              `json:"league_season_id"`
		SessionType                       int              `json:"session_type"`
		OrderID                           int              `json:"order_id"`
		MinLicenseLevel                   int              `json:"min_license_level"`
		MaxLicenseLevel                   int              `json:"max_license_level"`
		Status                            int              `json:"status"`
		PaceCarID                         interface{}      `json:"pace_car_id"`
		PaceCarClassID                    interface{}      `json:"pace_car_class_id"`
		NumOptLaps                        int              `json:"num_opt_laps"`
		DamageModel                       int              `json:"damage_model"`
		DoNotPaintCars                    bool             `json:"do_not_paint_cars"`
		GreenWhiteCheckeredLimit          int              `json:"green_white_checkered_limit"`
		DoNotCountCautionLaps             bool             `json:"do_not_count_caution_laps"`
		ConsecCautionsSingleFile          bool             `json:"consec_cautions_single_file"`
		ConsecCautionWithinNlaps          int              `json:"consec_caution_within_nlaps"`
		NoLapperWaveArounds               bool             `json:"no_lapper_wave_arounds"`
		ShortParadeLap                    bool             `json:"short_parade_lap"`
		StartOnQualTire                   bool             `json:"start_on_qual_tire"`
		TelemetryRestriction              int              `json:"telemetry_restriction"`
		TelemetryForceToDisk              int              `json:"telemetry_force_to_disk"`
		MaxAiDrivers                      int              `json:"max_ai_drivers"`
		AiAvoidPlayers                    bool             `json:"ai_avoid_players"`
		AdaptiveAiEnabled                 bool             `json:"adaptive_ai_enabled"`
		AdaptiveAiDifficulty              int              `json:"adaptive_ai_difficulty"`
		MustUseDiffTireTypesInRace        bool             `json:"must_use_diff_tire_types_in_race"`
		StartZone                         bool             `json:"start_zone"`
		EnablePitlaneCollisions           bool             `json:"enable_pitlane_collisions"`
		DisallowVirtualMirror             bool             `json:"disallow_virtual_mirror"`
		MaxVisorTearoffs                  int              `json:"max_visor_tearoffs"`
		CategoryID                        int              `json:"category_id"`
		Category                          string           `json:"category"`
		SessionFull                       bool             `json:"session_full"`
		Host                              CommonHost       `json:"host"`
		Track                             CommonTrack      `json:"track"`
		Weather                           CommonWeather    `json:"weather"`
		TrackState                        CommonTrackState `json:"track_state"`
		Farm                              CommonFarm       `json:"farm"`
		Admins                            []CommonAdmins   `json:"admins"`
		AllowedTeams                      []interface{}    `json:"allowed_teams"`
		AllowedLeagues                    []interface{}    `json:"allowed_leagues"`
		Cars                              []CommonCars     `json:"cars"`
		CountByCarID                      map[string]int   `json:"count_by_car_id"`
		CountByCarClassID                 map[string]int   `json:"count_by_car_class_id"`
		CarTypes                          []struct {
			CarType string `json:"car_type"`
		} `json:"car_types"`
		TrackTypes []struct {
			TrackType string `json:"track_type"`
		} `json:"track_types"`
		LicenseGroupTypes []struct {
			LicenseGroupType int `json:"license_group_type"`
		} `json:"license_group_types"`
		EventTypes []struct {
			EventType int `json:"event_type"`
		} `json:"event_types"`
		SessionTypes []struct {
			SessionType int `json:"session_type"`
		} `json:"session_types"`
		CanJoin        bool              `json:"can_join"`
		SessAdmin      bool              `json:"sess_admin"`
		Friends        []interface{}     `json:"friends"`
		Watched        []interface{}     `json:"watched"`
		EndTime        time.Time         `json:"end_time"`
		TeamEntryCount int               `json:"team_entry_count"`
		IsHeatRacing   bool              `json:"is_heat_racing"`
		Populated      bool              `json:"populated"`
		Broadcaster    bool              `json:"broadcaster"`
		MinIr          int               `json:"min_ir"`
		MaxIr          int               `json:"max_ir"`
		SessionDesc    string            `json:"session_desc,omitempty"`
		AiRosterName   string            `json:"ai_roster_name,omitempty"`
		HeatSesInfo    CommonHeatSesInfo `json:"heat_ses_info,omitempty"`
	} `json:"sessions"`
}

type HostedSessionsResp struct {
	Subscribed bool `json:"subscribed"`
	Sessions   []struct {
		AdaptiveAiDifficulty int            `json:"adaptive_ai_difficulty"`
		AdaptiveAiEnabled    bool           `json:"adaptive_ai_enabled"`
		Admins               []CommonAdmins `json:"admins"`
		AiAvoidPlayers       bool           `json:"ai_avoid_players"`
		AllowedLeagues       []interface{}  `json:"allowed_leagues"`
		AllowedTeams         []interface{}  `json:"allowed_teams"`
		CarTypes             []struct {
			CarType string `json:"car_type"`
		} `json:"car_types"`
		Cars                     []CommonCars   `json:"cars"`
		CarsLeft                 int            `json:"cars_left"`
		Category                 string         `json:"category"`
		CategoryID               int            `json:"category_id"`
		ConnectionBlackFlag      bool           `json:"connection_black_flag"`
		ConsecCautionWithinNlaps int            `json:"consec_caution_within_nlaps"`
		ConsecCautionsSingleFile bool           `json:"consec_cautions_single_file"`
		CountByCarClassID        map[string]int `json:"count_by_car_class_id"`
		CountByCarID             map[string]int `json:"count_by_car_id"`
		DamageModel              int            `json:"damage_model"`
		DisallowVirtualMirror    bool           `json:"disallow_virtual_mirror"`
		DoNotCountCautionLaps    bool           `json:"do_not_count_caution_laps"`
		DoNotPaintCars           bool           `json:"do_not_paint_cars"`
		DriverChangeRule         int            `json:"driver_change_rule"`
		DriverChanges            bool           `json:"driver_changes"`
		Elig                     CommonElig     `json:"elig"`
		EnablePitlaneCollisions  bool           `json:"enable_pitlane_collisions"`
		EntryCount               int            `json:"entry_count"`
		EventTypes               []struct {
			EventType int `json:"event_type"`
		} `json:"event_types"`
		Farm                     CommonFarm `json:"farm"`
		FixedSetup               bool       `json:"fixed_setup"`
		FullCourseCautions       bool       `json:"full_course_cautions"`
		GreenWhiteCheckeredLimit int        `json:"green_white_checkered_limit"`
		HardcoreLevel            int        `json:"hardcore_level"`
		Host                     CommonHost `json:"host"`
		IncidentLimit            int        `json:"incident_limit"`
		IncidentWarnMode         int        `json:"incident_warn_mode"`
		IncidentWarnParam1       int        `json:"incident_warn_param1"`
		IncidentWarnParam2       int        `json:"incident_warn_param2"`
		LaunchAt                 time.Time  `json:"launch_at"`
		LeagueID                 int        `json:"league_id"`
		LeagueSeasonID           int        `json:"league_season_id"`
		LicenseGroupTypes        []struct {
			LicenseGroupType int `json:"license_group_type"`
		} `json:"license_group_types"`
		LoneQualify                bool        `json:"lone_qualify"`
		LuckyDog                   bool        `json:"lucky_dog"`
		MaxAiDrivers               int         `json:"max_ai_drivers"`
		MaxDrivers                 int         `json:"max_drivers"`
		MaxIr                      int         `json:"max_ir"`
		MaxLicenseLevel            int         `json:"max_license_level"`
		MaxTeamDrivers             int         `json:"max_team_drivers"`
		MaxVisorTearoffs           int         `json:"max_visor_tearoffs"`
		MinIr                      int         `json:"min_ir"`
		MinLicenseLevel            int         `json:"min_license_level"`
		MinTeamDrivers             int         `json:"min_team_drivers"`
		MulticlassType             int         `json:"multiclass_type"`
		MustUseDiffTireTypesInRace bool        `json:"must_use_diff_tire_types_in_race"`
		NoLapperWaveArounds        bool        `json:"no_lapper_wave_arounds"`
		NumFastTows                int         `json:"num_fast_tows"`
		NumOptLaps                 int         `json:"num_opt_laps"`
		OpenRegExpires             time.Time   `json:"open_reg_expires"`
		OrderID                    int         `json:"order_id"`
		PaceCarClassID             interface{} `json:"pace_car_class_id"`
		PaceCarID                  interface{} `json:"pace_car_id"`
		PasswordProtected          bool        `json:"password_protected"`
		PitsInUse                  int         `json:"pits_in_use"`
		PracticeLength             int         `json:"practice_length"`
		PrivateSessionID           int         `json:"private_session_id"`
		QualifierMustStartRace     bool        `json:"qualifier_must_start_race"`
		QualifyLaps                int         `json:"qualify_laps"`
		QualifyLength              int         `json:"qualify_length"`
		RaceLaps                   int         `json:"race_laps"`
		RaceLength                 int         `json:"race_length"`
		Restarts                   int         `json:"restarts"`
		RestrictResults            bool        `json:"restrict_results"`
		RestrictViewing            bool        `json:"restrict_viewing"`
		RollingStarts              bool        `json:"rolling_starts"`
		SessionFull                bool        `json:"session_full"`
		SessionID                  int         `json:"session_id"`
		SessionName                string      `json:"session_name"`
		SessionType                int         `json:"session_type"`
		SessionTypes               []struct {
			SessionType int `json:"session_type"`
		} `json:"session_types"`
		ShortParadeLap       bool             `json:"short_parade_lap"`
		StartOnQualTire      bool             `json:"start_on_qual_tire"`
		StartZone            bool             `json:"start_zone"`
		Status               int              `json:"status"`
		SubsessionID         int              `json:"subsession_id"`
		TeamEntryCount       int              `json:"team_entry_count"`
		TelemetryForceToDisk int              `json:"telemetry_force_to_disk"`
		TelemetryRestriction int              `json:"telemetry_restriction"`
		TimeLimit            int              `json:"time_limit"`
		Track                CommonTrack      `json:"track"`
		TrackState           CommonTrackState `json:"track_state"`
		TrackTypes           []struct {
			TrackType string `json:"track_type"`
		} `json:"track_types"`
		UnsportConductRuleMode int               `json:"unsport_conduct_rule_mode"`
		WarmupLength           int               `json:"warmup_length"`
		Weather                CommonWeather     `json:"weather"`
		AiRosterName           string            `json:"ai_roster_name,omitempty"`
		HeatSesInfo            CommonHeatSesInfo `json:"heat_ses_info,omitempty"`
		SessionDesc            string            `json:"session_desc,omitempty"`
		RegisteredTeams        []int             `json:"registered_teams,omitempty"`
	} `json:"sessions"`
	Success bool `json:"success"`
}
