package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

// Position is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Position struct {
	gorm.Model
	Name string
}

// String is not required by pop and may be deleted
func (p Position) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Positions is not required by pop and may be deleted
type Positions []Position
