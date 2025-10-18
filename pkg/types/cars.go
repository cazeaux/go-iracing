package types

import "time"

type CarsGetReq struct{}

type CarsAssetsReq struct{}

type PaintRule struct {
	PaintCarAvailable       bool   `json:"PaintCarAvailable"`
	Color1                  string `json:"Color1"`
	Color2                  string `json:"Color2"`
	Color3                  string `json:"Color3"`
	Sponsor1Available       bool   `json:"Sponsor1Available"`
	Sponsor2Available       bool   `json:"Sponsor2Available"`
	Sponsor1                string `json:"Sponsor1"`
	Sponsor2                string `json:"Sponsor2"`
	RulesExplanation        string `json:"RulesExplanation"`
	PaintWheelAvailable     bool   `json:"PaintWheelAvailable,omitempty"`
	WheelColor              string `json:"WheelColor,omitempty"`
	RimTypeAvailable        bool   `json:"RimTypeAvailable,omitempty"`
	RimType                 string `json:"RimType,omitempty"`
	AllowNumberFontChanges  bool   `json:"AllowNumberFontChanges,omitempty"`
	NumberFont              string `json:"NumberFont,omitempty"`
	AllowNumberColorChanges bool   `json:"AllowNumberColorChanges,omitempty"`
	NumberColor1            string `json:"NumberColor1,omitempty"`
	NumberColor2            string `json:"NumberColor2,omitempty"`
	NumberColor3            string `json:"NumberColor3,omitempty"`
	RestrictCustomPaint     bool   `json:"RestrictCustomPaint,omitempty"`
}

type CarsGetResp struct {
	AiEnabled          bool   `json:"ai_enabled"`
	AllowNumberColors  bool   `json:"allow_number_colors"`
	AllowNumberFont    bool   `json:"allow_number_font"`
	AllowSponsor1      bool   `json:"allow_sponsor1"`
	AllowSponsor2      bool   `json:"allow_sponsor2"`
	AllowWheelColor    bool   `json:"allow_wheel_color"`
	AwardExempt        bool   `json:"award_exempt"`
	CarConfigDefs      []any  `json:"car_config_defs"`
	CarConfigs         []any  `json:"car_configs"`
	CarDirpath         string `json:"car_dirpath"`
	CarID              int    `json:"car_id"`
	CarName            string `json:"car_name"`
	CarNameAbbreviated string `json:"car_name_abbreviated"`
	CarTypes           []struct {
		CarType string `json:"car_type"`
	} `json:"car_types"`
	CarWeight               int       `json:"car_weight"`
	Categories              []string  `json:"categories"`
	Created                 time.Time `json:"created"`
	FirstSale               time.Time `json:"first_sale"`
	Folder                  string    `json:"folder"`
	ForumURL                string    `json:"forum_url,omitempty"`
	FreeWithSubscription    bool      `json:"free_with_subscription"`
	HasHeadlights           bool      `json:"has_headlights"`
	HasMultipleDryTireTypes bool      `json:"has_multiple_dry_tire_types"`
	HasRainCapableTireTypes bool      `json:"has_rain_capable_tire_types"`
	Hp                      int       `json:"hp"`
	IsPsPurchasable         bool      `json:"is_ps_purchasable"`
	Logo                    string    `json:"logo"`
	MaxPowerAdjustPct       int       `json:"max_power_adjust_pct"`
	MaxWeightPenaltyKg      int       `json:"max_weight_penalty_kg"`
	MinPowerAdjustPct       int       `json:"min_power_adjust_pct"`
	PackageID               int       `json:"package_id"`
	Patterns                int       `json:"patterns"`
	Price                   float64   `json:"price"`
	PriceDisplay            string    `json:"price_display,omitempty"`
	RainEnabled             bool      `json:"rain_enabled"`
	Retired                 bool      `json:"retired"`
	SearchFilters           string    `json:"search_filters"`
	Sku                     int       `json:"sku"`
	SmallImage              string    `json:"small_image"`
	SponsorLogo             any       `json:"sponsor_logo"`
	CarMake                 string    `json:"car_make,omitempty"`
	CarModel                string    `json:"car_model,omitempty"`
	PaintRules              *struct {
		Rule1               *PaintRule `json:"1"`
		Rule2               *PaintRule `json:"2"`
		Rule3               *PaintRule `json:"3"`
		Rule4               *PaintRule `json:"4"`
		Rule5               *PaintRule `json:"5"`
		Rule6               *PaintRule `json:"6"`
		Rule7               *PaintRule `json:"7"`
		Rule8               *PaintRule `json:"8"`
		Rule9               *PaintRule `json:"9"`
		Rule10              *PaintRule `json:"10"`
		RestrictCustomPaint bool       `json:"RestrictCustomPaint,omitempty"`
	} `json:"paint_rules"`
	SiteURL string `json:"site_url,omitempty"`
}

type CarsAssetsResp map[string]CarAsset

type CarAsset struct {
	CarID                  int    `json:"car_id"`
	CarRules               []any  `json:"car_rules"`
	DetailCopy             string `json:"detail_copy"`
	DetailScreenShotImages string `json:"detail_screen_shot_images"`
	DetailTechspecsCopy    string `json:"detail_techspecs_copy"`
	Folder                 string `json:"folder"`
	GalleryImages          string `json:"gallery_images"`
	GalleryPrefix          any    `json:"gallery_prefix"`
	GroupImage             any    `json:"group_image"`
	GroupName              any    `json:"group_name"`
	LargeImage             string `json:"large_image"`
	Logo                   string `json:"logo"`
	SmallImage             string `json:"small_image"`
	SponsorLogo            any    `json:"sponsor_logo"`
	TemplatePath           string `json:"template_path"`
}
