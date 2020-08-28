package model

import (
	"time"
)

type Task struct {
	Id        string
	Name      string
	Note 			string
	StartTime time.Time
	EndTime time.Time
	Status status
}

type status string

const TaskDone status = "DONE"
const TaskInProgress status = "IN_PROGRESS" 