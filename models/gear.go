package models

type Gear struct {
	Id          string  `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Name        string  `json:"name"`
	Primary     bool    `json:"primary"`
	Distance    float64 `json:"distance"`
	BrandName   string  `json:"brand_name"`
	ModelName   string  `json:"model_name"`
	FrameType   string  `json:"frame_type"`
	Description string  `json:"description"`
	AthleteID   uint
	Athlete     Athlete `gorm:"foreignkey:AthleteID" json:"athlete"`
}
