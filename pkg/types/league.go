package types

import "time"

// GET

type LeagueGetReq struct {
	LeagueID       int  `url:"league_id"`
	IncludeLicense bool `url:"include_license,omitempty"`
}

type LeagueGetResp struct {
	LeagueID        int       `json:"league_id"`
	OwnerID         int       `json:"owner_id"`
	LeagueName      string    `json:"league_name"`
	Created         time.Time `json:"created"`
	Hidden          bool      `json:"hidden"`
	Message         string    `json:"message"`
	About           string    `json:"about"`
	URL             string    `json:"url"`
	Recruiting      bool      `json:"recruiting"`
	Rules           string    `json:"rules"`
	PrivateWall     bool      `json:"private_wall"`
	PrivateRoster   bool      `json:"private_roster"`
	PrivateSchedule bool      `json:"private_schedule"`
	PrivateResults  bool      `json:"private_results"`
	IsOwner         bool      `json:"is_owner"`
	IsAdmin         bool      `json:"is_admin"`
	RosterCount     int       `json:"roster_count"`
	Owner           struct {
		CustID      int    `json:"cust_id"`
		DisplayName string `json:"display_name"`
		Helmet      struct {
			Pattern    int    `json:"pattern"`
			Color1     string `json:"color1"`
			Color2     string `json:"color2"`
			Color3     string `json:"color3"`
			FaceType   int    `json:"face_type"`
			HelmetType int    `json:"helmet_type"`
		} `json:"helmet"`
		CarNumber string `json:"car_number"`
		NickName  string `json:"nick_name"`
	} `json:"owner"`
	Image struct {
		SmallLogo string `json:"small_logo"`
		LargeLogo string `json:"large_logo"`
	} `json:"image"`
	Tags struct {
		Categorized    []string `json:"categorized"`
		NotCategorized []string `json:"not_categorized"`
	} `json:"tags"`
	LeagueApplications []string `json:"league_applications"`
	PendingRequests    []string `json:"pending_requests"`
	IsMember           bool     `json:"is_member"`
	IsApplicant        bool     `json:"is_applicant"`
	IsInvite           bool     `json:"is_invite"`
	IsIgnored          bool     `json:"is_ignored"`
	Roster             []struct {
		CustID      int    `json:"cust_id"`
		DisplayName string `json:"display_name"`
		Helmet      struct {
			Pattern    int    `json:"pattern"`
			Color1     string `json:"color1"`
			Color2     string `json:"color2"`
			Color3     string `json:"color3"`
			FaceType   int    `json:"face_type"`
			HelmetType int    `json:"helmet_type"`
		} `json:"helmet"`
		Owner             bool      `json:"owner"`
		Admin             bool      `json:"admin"`
		LeagueMailOptOut  bool      `json:"league_mail_opt_out"`
		LeaguePmOptOut    bool      `json:"league_pm_opt_out"`
		LeagueMemberSince time.Time `json:"league_member_since"`
		CarNumber         string    `json:"car_number"`
		NickName          string    `json:"nick_name"`
	} `json:"roster"`
}

// CustLeagueSessions

type LeagueCustLeagueSessionsReq struct {
	Mine      bool `url:"mine,omitempty"`
	PackageID int  `url:"package_id,omitempty"`
}

type LeagueCustLeagueSessionsResp struct {
	Mine       bool `json:"mine"`
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
		SessionDesc                       string           `json:"session_desc,omitempty"`
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
		LeagueName                        string           `json:"league_name"`
		LeagueSeasonID                    int              `json:"league_season_id"`
		LeagueSeasonName                  string           `json:"league_season_name,omitempty"`
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
		AdaptiveAiDifficulty              int              `json:"adaptive_ai_difficulty,omitempty"`
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
		AllowedLeagues                    []int            `json:"allowed_leagues"`
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
		CanJoin bool `json:"can_join"`
		Image   struct {
			SmallLogo string `json:"small_logo"`
			LargeLogo string `json:"large_logo"`
		} `json:"image"`
		Owner           bool              `json:"owner"`
		Admin           bool              `json:"admin"`
		Friends         []interface{}     `json:"friends"`
		Watched         []interface{}     `json:"watched"`
		EndTime         time.Time         `json:"end_time"`
		TeamEntryCount  int               `json:"team_entry_count"`
		IsHeatRacing    bool              `json:"is_heat_racing"`
		Populated       bool              `json:"populated"`
		Broadcaster     bool              `json:"broadcaster"`
		MinIr           int               `json:"min_ir"`
		MaxIr           int               `json:"max_ir"`
		AiRosterName    string            `json:"ai_roster_name,omitempty"`
		HeatSesInfo     CommonHeatSesInfo `json:"heat_ses_info,omitempty"`
		AiMinSkill      int               `json:"ai_min_skill,omitempty"`
		AiMaxSkill      int               `json:"ai_max_skill,omitempty"`
		RegisteredTeams []int             `json:"registered_teams,omitempty"`
	} `json:"sessions"`
	Success bool `json:"success"`
}

// Directory

type LeagueDirectoryReq struct {
	Search               string               `url:"search,omitempty"`
	Tag                  string               `url:"tag,omitempty"`
	RestrictToMember     bool                 `url:"restrict_to_member,omitempty"`
	RestrictToRecruiting bool                 `url:"restrict_to_recruiting,omitempty"`
	RestrictToFriends    bool                 `url:"restrict_to_friends,omitempty"`
	RestrictToWatched    bool                 `url:"restrict_to_watched,omitempty"`
	MinimumRosterCount   int                  `url:"minimum_roster_count,omitempty"`
	MaximumRosterCount   int                  `url:"maximum_roster_count,omitempty"`
	Lowerbound           int                  `url:"lowerbound,omitempty"`
	Upperbound           int                  `url:"upperbound,omitempty"`
	Sort                 LeagueDirectorySort  `url:"sort,omitempty"`  // relevance, leaguename, displayname, rostercount
	Order                LeagueDirectoryOrder `url:"order,omitempty"` // asc, desc
}

type LeagueDirectorySort string

const (
	LeagueDirectorySortRelevance   LeagueDirectorySort = "relevance"
	LeagueDirectorySortLeagueName  LeagueDirectorySort = "leaguename"
	LeagueDirectorySortOwnerName   LeagueDirectorySort = "displayname"
	LeagueDirectorySortRosterCount LeagueDirectorySort = "rostercount"
)

type LeagueDirectoryOrder string

const (
	LeagueDirectoryOrderAsc  LeagueDirectoryOrder = "asc"
	LeagueDirectoryOrderDesc LeagueDirectoryOrder = "desc"
)

type LeagueDirectoryResp struct {
	ResultsPage []struct {
		LeagueID           int         `json:"league_id"`
		OwnerID            int         `json:"owner_id"`
		LeagueName         string      `json:"league_name"`
		Created            time.Time   `json:"created"`
		About              string      `json:"about,omitempty"`
		URL                string      `json:"url,omitempty"`
		RosterCount        int         `json:"roster_count"`
		Recruiting         bool        `json:"recruiting"`
		IsAdmin            bool        `json:"is_admin"`
		IsMember           bool        `json:"is_member"`
		PendingApplication bool        `json:"pending_application"`
		PendingInvitation  bool        `json:"pending_invitation"`
		Owner              CommonOwner `json:"owner"`
	} `json:"results_page"`
	Success    bool `json:"success"`
	Lowerbound int  `json:"lowerbound"`
	Upperbound int  `json:"upperbound"`
	RowCount   int  `json:"row_count"`
}

// Get point systems

type LeagueGetPointsSystemsReq struct {
	LeagueID int `url:"league_id"`
	SeasonID int `url:"season_id,omitempty"`
}
type LeagueGetPointsSystemsResp struct {
	Subscribed    bool `json:"subscribed"`
	Success       bool `json:"success"`
	PointsSystems []struct {
		PointsSystemID int    `json:"points_system_id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		LeagueID       int    `json:"league_id"`
		Retired        bool   `json:"retired"`
		IracingSystem  bool   `json:"iracing_system"`
	} `json:"points_systems"`
	LeagueID int `json:"league_id"`
}

// Membership

type LeagueMembershipReq struct {
	CustID        int  `url:"cust_id,omitempty"`
	IncludeLeague bool `url:"include_league,omitempty"`
}

type LeagueMembershipResp []struct {
	LeagueID         int            `json:"league_id"`
	LeagueName       string         `json:"league_name"`
	Owner            bool           `json:"owner"`
	Admin            bool           `json:"admin"`
	LeagueMailOptOut bool           `json:"league_mail_opt_out"`
	LeaguePmOptOut   bool           `json:"league_pm_opt_out"`
	CarNumber        int            `json:"car_number"`
	NickName         string         `json:"nick_name"`
	League           *LeagueGetResp `json:"league"`
}

// Seasons

type LeagueSeasonsReq struct {
	LeagueID int  `url:"league_id,omitempty"`
	Retired  bool `url:"retired,omitempty"`
}

type LeagueSeasonsResp struct {
	Subscribed bool `json:"subscribed"`
	Seasons    []struct {
		LeagueID                int    `json:"league_id"`
		SeasonID                int    `json:"season_id"`
		PointsSystemID          int    `json:"points_system_id"`
		SeasonName              string `json:"season_name"`
		Active                  bool   `json:"active"`
		Hidden                  bool   `json:"hidden"`
		NumDrops                int    `json:"num_drops"`
		NoDropsOnOrAfterRaceNum int    `json:"no_drops_on_or_after_race_num"`
		PointsCars              []struct {
			CarID   int    `json:"car_id"`
			CarName string `json:"car_name"`
		} `json:"points_cars"`
		DriverPointsCarClasses []struct {
			CarClassID  int    `json:"car_class_id"`
			Name        string `json:"name"`
			CarsInClass []struct {
				CarID   int    `json:"car_id"`
				CarName string `json:"car_name"`
			} `json:"cars_in_class"`
		} `json:"driver_points_car_classes"`
		TeamPointsCarClasses []struct {
			CarClassID  int    `json:"car_class_id"`
			Name        string `json:"name"`
			CarsInClass []struct {
				CarID   int    `json:"car_id"`
				CarName string `json:"car_name"`
			} `json:"cars_in_class"`
		} `json:"team_points_car_classes"`
		PointsSystemName string `json:"points_system_name"`
		PointsSystemDesc string `json:"points_system_desc"`
	} `json:"seasons"`
	Success  bool `json:"success"`
	Retired  bool `json:"retired"`
	LeagueID int  `json:"league_id"`
}

// Season standings

type LeagueSeasonStandingsReq struct {
	LeagueID int `url:"league_id"`
	SeasonID int `url:"season_id"`
}

type LeagueSeasonStandingsResp struct {
	CarClassID int  `json:"car_class_id"`
	Success    bool `json:"success"`
	SeasonID   int  `json:"season_id"`
	CarID      int  `json:"car_id"`
	Standings  struct {
		DriverStandings []struct {
			Rownum   int `json:"rownum"`
			Position int `json:"position"`
			Driver   struct {
				CustID      int    `json:"cust_id"`
				DisplayName string `json:"display_name"`
				Helmet      struct {
					Pattern    int    `json:"pattern"`
					Color1     string `json:"color1"`
					Color2     string `json:"color2"`
					Color3     string `json:"color3"`
					FaceType   int    `json:"face_type"`
					HelmetType int    `json:"helmet_type"`
				} `json:"helmet"`
			} `json:"driver"`
			CarNumber           string `json:"car_number"`
			DriverNickname      string `json:"driver_nickname"`
			Wins                int    `json:"wins"`
			AverageStart        int    `json:"average_start"`
			AverageFinish       int    `json:"average_finish"`
			BasePoints          int    `json:"base_points"`
			NegativeAdjustments int    `json:"negative_adjustments"`
			PositiveAdjustments int    `json:"positive_adjustments"`
			TotalAdjustments    int    `json:"total_adjustments"`
			TotalPoints         int    `json:"total_points"`
		} `json:"driver_standings"`
		TeamStandings         []interface{} `json:"team_standings"`
		DriverStandingsCsvURL string        `json:"driver_standings_csv_url"`
		TeamStandingsCsvURL   string        `json:"team_standings_csv_url"`
	} `json:"standings"`
	LeagueID int `json:"league_id"`
}

// Season sessions

type LeagueSeasonSessionsReq struct {
	LeagueID    int  `url:"league_id"`
	SeasonID    int  `url:"season_id"`
	ResultsOnly bool `url:"results_only,omitempty"`
}
type LeagueSeasonSessionsResp struct {
	Success    bool `json:"success"`
	Subscribed bool `json:"subscribed"`
	LeagueID   int  `json:"league_id"`
	SeasonID   int  `json:"season_id"`
	Sessions   []struct {
		Cars []struct {
			CarID        int    `json:"car_id"`
			CarName      string `json:"car_name"`
			CarClassID   int    `json:"car_class_id"`
			CarClassName string `json:"car_class_name"`
		} `json:"cars"`
		DriverChanges     bool             `json:"driver_changes"`
		EntryCount        int              `json:"entry_count"`
		HasResults        bool             `json:"has_results"`
		LaunchAt          time.Time        `json:"launch_at"`
		LeagueID          int              `json:"league_id"`
		LeagueSeasonID    int              `json:"league_season_id"`
		LoneQualify       bool             `json:"lone_qualify"`
		PaceCarClassID    int              `json:"pace_car_class_id"`
		PaceCarID         int              `json:"pace_car_id"`
		PasswordProtected bool             `json:"password_protected"`
		PracticeLength    int              `json:"practice_length"`
		PrivateSessionID  int              `json:"private_session_id"`
		QualifyLaps       int              `json:"qualify_laps"`
		QualifyLength     int              `json:"qualify_length"`
		RaceLaps          int              `json:"race_laps"`
		RaceLength        int              `json:"race_length"`
		SessionID         int              `json:"session_id"`
		Status            int              `json:"status"`
		SubsessionID      int              `json:"subsession_id"`
		TeamEntryCount    int              `json:"team_entry_count"`
		TimeLimit         int              `json:"time_limit"`
		Track             CommonTrack      `json:"track"`
		TrackState        CommonTrackState `json:"track_state"`
		Weather           CommonWeather    `json:"weather"`
		WinnerID          int              `json:"winner_id"`
		WinnerName        string           `json:"winner_name"`
	} `json:"sessions"`
}
