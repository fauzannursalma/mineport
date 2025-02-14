package entity

import "gorm.io/gorm"

type NonFormalEducation struct {
	gorm.Model
	UserID      uint   `gorm:"not null"`
	Provider       string
	CourseName     string
	CompletionYear int
	Image          string

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}