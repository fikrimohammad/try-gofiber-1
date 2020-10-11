package models

import "gorm.io/gorm"

// Club is a model to symbolize record on clubs table.
type Club struct {
	gorm.Model
	FullName     string
	ShortName    string
	Abbreviation string
	IconURL      string
	WebURL       string
	Stadium      Stadium
}

// Clubs is a model to symbolize array of Club
type Clubs []Club
