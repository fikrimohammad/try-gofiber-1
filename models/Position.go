package models

import (
	"gorm.io/gorm"
)

// Position is a model to symbolize record on positions table.
type Position struct {
	gorm.Model
	Name string
}

// Positions is a model to symbolize array of Position
type Positions []Position
