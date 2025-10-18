package types

type CarClassGetReq struct{}

type CarClassGetResp struct {
	CarClassID    int           `json:"car_class_id"`
	CarsInClass   []CarsInClass `json:"cars_in_class"`
	CustID        int           `json:"cust_id"`
	Name          string        `json:"name"`
	RainEnabled   bool          `json:"rain_enabled"`
	RelativeSpeed int           `json:"relative_speed"`
	ShortName     string        `json:"short_name"`
}
type CarsInClass struct {
	CarDirpath  string `json:"car_dirpath"`
	CarID       int    `json:"car_id"`
	RainEnabled bool   `json:"rain_enabled"`
	Retired     bool   `json:"retired"`
}
