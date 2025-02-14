package entity

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Name        string
	Level string 			// Beginner, Intermediate, Advanced
	Image       string

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

}