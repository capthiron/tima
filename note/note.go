package note

import (
	"gorm.io/gorm"
	"time"
)

type Note struct {
	gorm.Model
	Description string
	Start       time.Time
	End         time.Time
	Status      Status
}

type Status string

const Done Status = "DONE"
const InProgress Status = "IN_PROGRESS"
