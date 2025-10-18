package types

// Licenses

type LookupLicensesReq struct {
}

type LookupLicensesResp struct {
	LicenseGroup         int    `json:"license_group"`
	GroupName            string `json:"group_name"`
	MinNumRaces          int    `json:"min_num_races"`
	ParticipationCredits int    `json:"participation_credits"`
	MinSrToFastTrack     int    `json:"min_sr_to_fast_track"`
	Levels               []struct {
		LicenseID     int    `json:"license_id"`
		LicenseGroup  int    `json:"license_group"`
		License       string `json:"license"`
		ShortName     string `json:"short_name"`
		LicenseLetter string `json:"license_letter"`
		Color         string `json:"color"`
	} `json:"levels"`
	MinNumTt int `json:"min_num_tt"`
}

// Countries

type LookupCountriesReq struct {
}

type LookupCountriesResp struct {
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`
}

// Flairs

type LookupFlairsReq struct {
}

type LookupFlairsResp struct {
	FlairID        int    `json:"flair_id"`
	FlairName      string `json:"flair_name"`
	FlairShortname string `json:"flair_shortname"`
	CountryCode    string `json:"country_code"`
	Seq            int    `json:"seq"`
}

// Drivers

type LookupDriversReq struct {
	SearchTerm string `url:"search_term"`
	LeagueID   int    `url:"league_id"`
}
type LookupDriversResp struct {
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
	ProfileDisabled bool `json:"profile_disabled"`
}
