package types

import "time"

type MemberInfoResponse struct {
	CustID           int       `json:"cust_id"`
	DisplayName      string    `json:"display_name"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	OnCarName        string    `json:"on_car_name"`
	MemberSince      string    `json:"member_since"`
	FlairID          int       `json:"flair_id"`
	FlairName        string    `json:"flair_name"`
	FlairShortname   string    `json:"flair_shortname"`
	FlairCountryCode string    `json:"flair_country_code"`
	LastLogin        time.Time `json:"last_login"`
	ReadTc           time.Time `json:"read_tc"`
	ReadPp           time.Time `json:"read_pp"`
	ReadCompRules    time.Time `json:"read_comp_rules"`
	Flags            int       `json:"flags"`
	ConnectionType   string    `json:"connection_type"`
	DownloadServer   string    `json:"download_server"`
	Account          struct {
		IrDollars    int    `json:"ir_dollars"`
		IrCredits    int    `json:"ir_credits"`
		Status       string `json:"status"`
		CountryRules any    `json:"country_rules"`
	} `json:"account"`
	Helmet CommonHelmet `json:"helmet"`
	Suit   struct {
		Pattern  int    `json:"pattern"`
		Color1   string `json:"color1"`
		Color2   string `json:"color2"`
		Color3   string `json:"color3"`
		BodyType int    `json:"body_type"`
	} `json:"suit"`
	Licenses struct {
		Oval       CommonMemberLicense `json:"oval"`
		SportsCar  CommonMemberLicense `json:"sports_car"`
		FormulaCar CommonMemberLicense `json:"formula_car"`
		DirtOval   CommonMemberLicense `json:"dirt_oval"`
		DirtRoad   CommonMemberLicense `json:"dirt_road"`
	} `json:"licenses"`
	CarPackages []struct {
		PackageID  int   `json:"package_id"`
		ContentIds []int `json:"content_ids"`
	} `json:"car_packages"`
	TrackPackages []struct {
		PackageID  int   `json:"package_id"`
		ContentIds []int `json:"content_ids"`
	} `json:"track_packages"`
	OtherOwnedPackages []int `json:"other_owned_packages"`
	Dev                bool  `json:"dev"`
	AlphaTester        bool  `json:"alpha_tester"`
	RainTester         bool  `json:"rain_tester"`
	Broadcaster        bool  `json:"broadcaster"`
	Restrictions       struct {
	} `json:"restrictions"`
	HasReadCompRules     bool   `json:"has_read_comp_rules"`
	HasReadNda           bool   `json:"has_read_nda"`
	FlagsHex             string `json:"flags_hex"`
	HundredPctClub       bool   `json:"hundred_pct_club"`
	TwentyPctDiscount    bool   `json:"twenty_pct_discount"`
	LastSeason           int    `json:"last_season"`
	HasAdditionalContent bool   `json:"has_additional_content"`
	HasReadTc            bool   `json:"has_read_tc"`
	HasReadPp            bool   `json:"has_read_pp"`
}

// data/member/get

type MemberGetReq struct {
	CustIDs         []int `url:"cust_ids,comma"`
	IncludeLicenses bool  `url:"include_licenses,omitempty"`
}

type MemberGetResp struct {
	Success bool  `json:"success"`
	CustIds []int `json:"cust_ids"`
	Members []struct {
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
		LastLogin      time.Time `json:"last_login"`
		MemberSince    string    `json:"member_since"`
		FlairID        int       `json:"flair_id"`
		FlairName      string    `json:"flair_name"`
		FlairShortname string    `json:"flair_shortname"`
		Ai             bool      `json:"ai"`
	} `json:"members"`
}

// Awards

type MemberAwardsReq struct {
	CustID int `url:"cust_id,omitempty"`
}

type MemberAwardsResp struct {
	MemberAwardID       int    `json:"member_award_id"`
	AwardID             int    `json:"award_id"`
	Achievement         bool   `json:"achievement"`
	AwardCount          int    `json:"award_count"`
	AwardOrder          int    `json:"award_order"`
	CustID              int    `json:"cust_id"`
	Description         string `json:"description,omitempty"`
	GroupName           string `json:"group_name"`
	HasPdf              bool   `json:"has_pdf"`
	IconBackgroundColor string `json:"icon_background_color,omitempty"`
	IconURLLarge        string `json:"icon_url_large"`
	IconURLSmall        string `json:"icon_url_small"`
	IconURLUnawarded    string `json:"icon_url_unawarded"`
	Name                string `json:"name"`
	Weight              int    `json:"weight"`
	AwardDate           string `json:"award_date,omitempty"`
	AwardedDescription  string `json:"awarded_description,omitempty"`
	DisplayDate         string `json:"display_date,omitempty"`
	SubsessionID        int    `json:"subsession_id,omitempty"`
	Viewed              bool   `json:"viewed,omitempty"`
	Progress            int    `json:"progress,omitempty"`
	ProgressLabel       string `json:"progress_label,omitempty"`
	Threshold           int    `json:"threshold,omitempty"`
	ProgressText        string `json:"progress_text,omitempty"`
	ProgressTextLabel   string `json:"progress_text_label,omitempty"`
}

// Award Instances

type MemberAwardsInstancesReq struct {
	CustID  int `url:"cust_id,omitempty"`
	AwardID int `url:"award_id"`
}

type MemberAwardsInstancesResp struct {
	CustID      int  `json:"cust_id"`
	AwardID     int  `json:"award_id"`
	Achievement bool `json:"achievement"`
	AwardCount  int  `json:"award_count"`
	Awards      []struct {
		MemberAwardID      int    `json:"member_award_id"`
		AwardID            int    `json:"award_id"`
		Achievement        bool   `json:"achievement"`
		AwardCount         int    `json:"award_count"`
		AwardDate          string `json:"award_date"`
		AwardOrder         int    `json:"award_order"`
		AwardedDescription string `json:"awarded_description"`
		CustID             int    `json:"cust_id"`
		Progress           int    `json:"progress"`
		ProgressLabel      string `json:"progress_label"`
		Threshold          int    `json:"threshold"`
		Viewed             bool   `json:"viewed"`
	} `json:"awards"`
	Description      string `json:"description"`
	GroupName        string `json:"group_name"`
	HasPdf           bool   `json:"has_pdf"`
	IconURLLarge     string `json:"icon_url_large"`
	IconURLSmall     string `json:"icon_url_small"`
	IconURLUnawarded string `json:"icon_url_unawarded"`
	Name             string `json:"name"`
	Weight           int    `json:"weight"`
}

// Chart Data

type MemberChartDataReq struct {
	CustID     int `url:"cust_id,omitempty"`
	ChartType  int `url:"chart_type"`  //1 - iRating; 2 - TT Rating; 3 - License/SR
	CategoryID int `url:"category_id"` //1 - Oval; 2 - Road; 3 - Dirt oval; 4 - Dirt road : 5 - Sports Car - 6 : Formula Car
}

type MemberChartDataResp struct {
	Blackout   bool `json:"blackout"`
	CategoryID int  `json:"category_id"`
	ChartType  int  `json:"chart_type"`
	Data       []struct {
		When  string `json:"when"`
		Value int    `json:"value"`
	} `json:"data"`
	Success bool `json:"success"`
	CustID  int  `json:"cust_id"`
}

// Profile

type MemberProfileReq struct {
	CustID int `url:"cust_id,omitempty"`
}

type MemberProfileResp struct {
	RecentAwards []struct {
		MemberAwardID      int    `json:"member_award_id"`
		AwardID            int    `json:"award_id"`
		Achievement        bool   `json:"achievement"`
		AwardCount         int    `json:"award_count"`
		AwardDate          string `json:"award_date"`
		AwardOrder         int    `json:"award_order"`
		AwardedDescription string `json:"awarded_description"`
		CustID             int    `json:"cust_id"`
		Description        string `json:"description"`
		GroupName          string `json:"group_name"`
		HasPdf             bool   `json:"has_pdf"`
		IconURLLarge       string `json:"icon_url_large"`
		IconURLSmall       string `json:"icon_url_small"`
		IconURLUnawarded   string `json:"icon_url_unawarded"`
		Name               string `json:"name"`
		Viewed             bool   `json:"viewed"`
		Weight             int    `json:"weight"`
		SubsessionID       int    `json:"subsession_id,omitempty"`
	} `json:"recent_awards"`
	Activity struct {
		Recent30DaysCount    int `json:"recent_30days_count"`
		Prev30DaysCount      int `json:"prev_30days_count"`
		ConsecutiveWeeks     int `json:"consecutive_weeks"`
		MostConsecutiveWeeks int `json:"most_consecutive_weeks"`
	} `json:"activity"`
	Success    bool   `json:"success"`
	ImageURL   string `json:"image_url"`
	MemberInfo struct {
		Ai             bool                  `json:"ai"`
		Country        string                `json:"country"`
		CountryCode    string                `json:"country_code"`
		CustID         int                   `json:"cust_id"`
		DisplayName    string                `json:"display_name"`
		FlairID        int                   `json:"flair_id"`
		FlairName      string                `json:"flair_name"`
		FlairShortname string                `json:"flair_shortname"`
		Helmet         CommonHelmet          `json:"helmet"`
		LastLogin      time.Time             `json:"last_login"`
		Licenses       []CommonMemberLicense `json:"licenses"`
		MemberSince    string                `json:"member_since"`
	} `json:"member_info"`
	Disabled       bool                  `json:"disabled"`
	LicenseHistory []CommonMemberLicense `json:"license_history"`
	RecentEvents   []struct {
		EventType        string      `json:"event_type"`
		SubsessionID     int         `json:"subsession_id"`
		StartTime        time.Time   `json:"start_time"`
		EventID          int         `json:"event_id"`
		EventName        string      `json:"event_name"`
		SimsessionType   int         `json:"simsession_type"`
		StartingPosition int         `json:"starting_position"`
		FinishPosition   int         `json:"finish_position"`
		BestLapTime      int         `json:"best_lap_time"`
		PercentRank      int         `json:"percent_rank"`
		CarID            int         `json:"car_id"`
		CarName          string      `json:"car_name"`
		LogoURL          string      `json:"logo_url"`
		Track            CommonTrack `json:"track"`
	} `json:"recent_events"`
	CustID         int  `json:"cust_id"`
	IsGenericImage bool `json:"is_generic_image"`
	FollowCounts   struct {
		Followers int `json:"followers"`
		Follows   int `json:"follows"`
	} `json:"follow_counts"`
}

// Participation Credits

type MemberParticipationCreditsReq struct{}

type MemberParticipationCreditsResp struct {
	CustID               int     `json:"cust_id"`
	SeasonID             int     `json:"season_id"`
	SeriesID             int     `json:"series_id"`
	SeriesName           string  `json:"series_name"`
	LicenseGroup         int     `json:"license_group"`
	LicenseGroupName     string  `json:"license_group_name"`
	ParticipationCredits float64 `json:"participation_credits"`
	MinWeeks             int     `json:"min_weeks"`
	Weeks                int     `json:"weeks"`
	EarnedCredits        float64 `json:"earned_credits"`
	TotalCredits         float64 `json:"total_credits"`
}
