package main

type CueSchema struct {
	ID      int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Cuelang string `gorm:"not null" form:"Cuelang" json:"Cuelang"`
}
