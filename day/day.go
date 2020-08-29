package day

import (
	"github.com/capthiron/tima/task"
	"gorm.io/gorm"
	"time"
)

type Day struct {
	gorm.Model
	StartTime     time.Time `gorm:"not null"`
	EndTime       time.Time
	BreakDuration int
	Tasks         []task.Task `gorm:"-"`
}
