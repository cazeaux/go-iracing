package types

import "time"

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
		CarNumber any `json:"car_number"`
		NickName  any `json:"nick_name"`
	} `json:"owner"`
	Image struct {
		SmallLogo any `json:"small_logo"`
		LargeLogo any `json:"large_logo"`
	} `json:"image"`
	Tags struct {
		Categorized    []any `json:"categorized"`
		NotCategorized []any `json:"not_categorized"`
	} `json:"tags"`
	LeagueApplications []any `json:"league_applications"`
	PendingRequests    []any `json:"pending_requests"`
	IsMember           bool  `json:"is_member"`
	IsApplicant        bool  `json:"is_applicant"`
	IsInvite           bool  `json:"is_invite"`
	IsIgnored          bool  `json:"is_ignored"`
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
		NickName          any       `json:"nick_name"`
	} `json:"roster"`
}
