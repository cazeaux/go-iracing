package types

import "time"

type CommonElig struct {
	SessionFull     bool  `json:"session_full"`
	CanSpot         bool  `json:"can_spot"`
	CanWatch        bool  `json:"can_watch"`
	CanDrive        bool  `json:"can_drive"`
	HasSessPassword bool  `json:"has_sess_password"`
	NeedsPurchase   bool  `json:"needs_purchase"`
	OwnCar          bool  `json:"own_car"`
	OwnTrack        bool  `json:"own_track"`
	PurchaseSkus    []int `json:"purchase_skus"`
	Registered      bool  `json:"registered"`
}

type CommonAdmins struct {
	CustID      int          `json:"cust_id"`
	DisplayName string       `json:"display_name"`
	Helmet      CommonHelmet `json:"helmet"`
}

type CommonOwner struct {
	CustID      int          `json:"cust_id"`
	DisplayName string       `json:"display_name"`
	Helmet      CommonHelmet `json:"helmet"`
	CarNumber   interface{}  `json:"car_number"`
	NickName    interface{}  `json:"nick_name"`
}

type CommonCars struct {
	CarID             int     `json:"car_id"`
	CarName           string  `json:"car_name"`
	CarClassID        int     `json:"car_class_id"`
	CarClassName      string  `json:"car_class_name"`
	MaxPctFuelFill    float64 `json:"max_pct_fuel_fill"`
	WeightPenaltyKg   float64 `json:"weight_penalty_kg"`
	PowerAdjustPct    float64 `json:"power_adjust_pct"`
	MaxDryTireSets    int     `json:"max_dry_tire_sets"`
	QualSetupID       int     `json:"qual_setup_id"`
	QualSetupFilename string  `json:"qual_setup_filename"`
	RaceSetupID       int     `json:"race_setup_id"`
	RaceSetupFilename string  `json:"race_setup_filename"`
	PackageID         int     `json:"package_id"`
}

type CommonFarm struct {
	FarmID      int    `json:"farm_id"`
	DisplayName string `json:"display_name"`
	ImagePath   string `json:"image_path"`
	Displayed   bool   `json:"displayed"`
}

type CommonHost struct {
	CustID      int          `json:"cust_id"`
	DisplayName string       `json:"display_name"`
	Helmet      CommonHelmet `json:"helmet"`
}

type CommonHelmet struct {
	Pattern    int    `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int    `json:"face_type"`
	HelmetType int    `json:"helmet_type"`
}

type CommonTrack struct {
	CategoryID int    `json:"category_id"`
	TrackID    int    `json:"track_id"`
	TrackName  string `json:"track_name"`
	ConfigName string `json:"config_name"`
}

type CommonTrackState struct {
	LeaveMarbles         bool `json:"leave_marbles"`
	PracticeGripCompound int  `json:"practice_grip_compound"`
	PracticeRubber       int  `json:"practice_rubber"`
	QualifyGripCompound  int  `json:"qualify_grip_compound"`
	QualifyRubber        int  `json:"qualify_rubber"`
	RaceGripCompound     int  `json:"race_grip_compound"`
	RaceRubber           int  `json:"race_rubber"`
	WarmupGripCompound   int  `json:"warmup_grip_compound"`
	WarmupRubber         int  `json:"warmup_rubber"`
}

type CommonWeather struct {
	AllowFog                bool    `json:"allow_fog"`
	Fog                     int     `json:"fog"`
	PrecipOption            int     `json:"precip_option"`
	RelHumidity             float64 `json:"rel_humidity"`
	SimulatedStartTime      string  `json:"simulated_start_time"`
	SimulatedTimeMultiplier int     `json:"simulated_time_multiplier"`
	SimulatedTimeOffsets    []int   `json:"simulated_time_offsets"`
	Skies                   int     `json:"skies"`
	TempUnits               int     `json:"temp_units"`
	TempValue               float64 `json:"temp_value"`
	TimeOfDay               int     `json:"time_of_day"`
	TrackWater              float64 `json:"track_water"`
	Type                    int     `json:"type"`
	Version                 int     `json:"version"`
	WeatherURL              string  `json:"weather_url"`
	WindDir                 float64 `json:"wind_dir"`
	WindUnits               int     `json:"wind_units"`
	WindValue               float64 `json:"wind_value"`
}

type CommonHeatSesInfo struct {
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
}

type CommonMemberLicense struct {
	CategoryID    int     `json:"category_id"`
	Category      string  `json:"category"`
	CategoryName  string  `json:"category_name"`
	LicenseLevel  int     `json:"license_level"`
	SafetyRating  float64 `json:"safety_rating"`
	Cpi           float64 `json:"cpi"`
	Irating       int     `json:"irating"`
	TtRating      int     `json:"tt_rating"`
	MprNumRaces   int     `json:"mpr_num_races"`
	Color         string  `json:"color"`
	GroupName     string  `json:"group_name"`
	GroupID       int     `json:"group_id"`
	ProPromotable bool    `json:"pro_promotable"`
	Seq           int     `json:"seq"`
	MprNumTts     int     `json:"mpr_num_tts"`
}
