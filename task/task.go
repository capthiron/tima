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
	Status    Status
}

type Status string

const TaskDone Status = "DONE"
const TaskInProgress Status = "IN_PROGRESS"
