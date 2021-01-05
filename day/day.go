package day

import (
	"github.com/capthiron/tima/note"
	"gorm.io/gorm"
	"time"
)

type Day struct {
	gorm.Model
	StartTime     time.Time `gorm:"not null"`
	EndTime       time.Time
	BreakDuration int
	WorkHours     WorkHours   `gorm:"-"`
	Notes         []note.Note `gorm:"-"`
}

type WorkHours struct {
	hours   int
	minutes int
}
