package models

import (
	"time"

	"gorm.io/gorm"
)

// Stadium is a model to symbolize record on stadiums table.
type Stadium struct {
	gorm.Model
	Name        string
	Description string
	Capacity    int
	BuiltOn     time.Time
	ClubID      int
}
