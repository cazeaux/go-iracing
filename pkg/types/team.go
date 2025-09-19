package types

import "time"

// Get

type TeamGetReq struct {
	TeamID          int `url:"team_id"`
	IncludeLicenses int `url:"include_licenses,omitempty"`
}

type TeamGetResp struct {
	TeamID      int       `json:"team_id"`
	OwnerID     int       `json:"owner_id"`
	TeamName    string    `json:"team_name"`
	Created     time.Time `json:"created"`
	Hidden      bool      `json:"hidden"`
	About       string    `json:"about"`
	URL         string    `json:"url"`
	RosterCount int       `json:"roster_count"`
	Recruiting  bool      `json:"recruiting"`
	PrivateWall bool      `json:"private_wall"`
	IsDefault   bool      `json:"is_default"`
	IsOwner     bool      `json:"is_owner"`
	IsAdmin     bool      `json:"is_admin"`
	Suit        struct {
		Pattern int    `json:"pattern"`
		Color1  string `json:"color1"`
		Color2  string `json:"color2"`
		Color3  string `json:"color3"`
	} `json:"suit"`
	Owner struct {
		CustID      int          `json:"cust_id"`
		DisplayName string       `json:"display_name"`
		Helmet      CommonHelmet `json:"helmet"`
		Owner       bool         `json:"owner"`
		Admin       bool         `json:"admin"`
	} `json:"owner"`
	Tags struct {
		Categorized    []interface{} `json:"categorized"`
		NotCategorized []interface{} `json:"not_categorized"`
	} `json:"tags"`
	TeamApplications []interface{} `json:"team_applications"`
	PendingRequests  []interface{} `json:"pending_requests"`
	IsMember         bool          `json:"is_member"`
	IsApplicant      bool          `json:"is_applicant"`
	IsInvite         bool          `json:"is_invite"`
	IsIgnored        bool          `json:"is_ignored"`
	Roster           []struct {
		CustID      int                   `json:"cust_id"`
		DisplayName string                `json:"display_name"`
		Helmet      CommonHelmet          `json:"helmet"`
		Licenses    []CommonMemberLicense `json:"licenses"`
		Owner       bool                  `json:"owner"`
		Admin       bool                  `json:"admin"`
	} `json:"roster"`
}

// Membership

type TeamMembershipReq struct{}

type TeamMembershipResp struct {
	TeamID      int    `json:"team_id"`
	TeamName    string `json:"team_name"`
	Owner       bool   `json:"owner"`
	Admin       bool   `json:"admin"`
	DefaultTeam bool   `json:"default_team"`
}
