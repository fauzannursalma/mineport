package config

import (
	"fmt"

	"github.com/fauzannursalma/mineport/internal/entity"
)

func MigrateDB() {
    if DB == nil {
        panic("Database is not connected")
    }

    err := DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
        &entity.User{},
        &entity.FormalEducation{},
        &entity.NonFormalEducation{},
        &entity.WorkExperience{},
        &entity.Skill{},
        &entity.Project{},
        &entity.Certificate{},
    )

    if err != nil {
        panic("Migration failed: " + err.Error())
    }

    fmt.Println("Database migrated!")
}
