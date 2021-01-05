package config

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	DayLength            float64
	DefaultBreakDuration int
	DateFormat           DateFormat `gorm:"default"`
}

type DateFormat string
