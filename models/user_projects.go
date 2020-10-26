package models

import (
	"time"
)

// @docs https://gorm.io/docs/connecting_to_the_database.html
type UserProjects struct {
	UserProjectId uint32 `gorm:"primary_key"`
	Uid uint32
	Vid uint32
	ProjectName string
	ProjectToken string
	Type uint8
	Active uint8
	CreatedAt time.Time
	DeletedAt time.Time
}

