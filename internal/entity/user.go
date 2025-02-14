package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(255); not null"`
	Email string `gorm:"type:varchar(255); not null; unique"`
	Password string `gorm:"type:varchar(255); not null"`
	ProfileImage string 

	FormalEducations   []FormalEducation     `gorm:"foreignKey:UserID"`
	NonFormalEducations []NonFormalEducation `gorm:"foreignKey:UserID"`
	WorkExperiences    []WorkExperience      `gorm:"foreignKey:UserID"`
	Skills            []Skill                `gorm:"foreignKey:UserID"`
	Projects          []Project              `gorm:"foreignKey:UserID"`
	Certificates      []Certificate          `gorm:"foreignKey:UserID"`
}