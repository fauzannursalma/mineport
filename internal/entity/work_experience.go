package entity

import "gorm.io/gorm"

type WorkExperience struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	CompanyName string
	Position    string
	StartDate   string // Format: "YYYY-MM"
	EndDate     string // Format: "YYYY-MM"
	Description string

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
