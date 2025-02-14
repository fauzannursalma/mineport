package entity

import "gorm.io/gorm"

type FormalEducation struct {
	gorm.Model
	UserID      uint   `gorm:"not null"`
	Institution  string
	Degree       string
	FieldOfStudy string
	StartYear    int
	EndYear      int

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
