package entity

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	UserID      uint   `gorm:"not null"`
	Title       string
	Description string
	TechStack   string
	Link        string
	Image       string

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

}