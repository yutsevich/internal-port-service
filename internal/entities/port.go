package entities

type Port struct {
	ID          uint          `json:"id,omitempty"`
	Name        string        `json:"name,omitempty"`
	City        string        `json:"city,omitempty"`
	Country     string        `json:"country,omitempty"`
	Alias       []interface{} `gorm:"type:text" json:"alias,omitempty"`
	Regions     []interface{} `gorm:"type:text" json:"regions,omitempty"`
	Coordinates []float64     `gorm:"type:text" json:"coordinates,omitempty"`
	Province    string        `json:"province,omitempty"`
	Timezone    string        `json:"timezone,omitempty"`
	Unlocs      []string      `gorm:"type:text" json:"unlocs,omitempty"`
	Code        string        `json:"code,omitempty"`
}

//type T struct {
//	AEAJM struct {
//		Name        string        `json:"name"`
//		City        string        `json:"city"`
//		Country     string        `json:"country"`
//		Alias       []interface{} `json:"alias"`
//		Regions     []interface{} `json:"regions"`
//		Coordinates []float64     `json:"coordinates"`
//		Province    string        `json:"province"`
//		Timezone    string        `json:"timezone"`
//		Unlocs      []string      `json:"unlocs"`
//		Code        string        `json:"code"`
//	} `json:"AEAJM"`
//}
