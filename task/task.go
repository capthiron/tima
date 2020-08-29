package task

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Name      string
	Note      string
	StartTime time.Time
	EndTime   time.Time
	Status    status
}

type status string

const TaskDone status = "DONE"
const TaskInProgress status = "IN_PROGRESS"
