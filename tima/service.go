package tima

import (
	"github.com/capthiron/tima/day"
	"github.com/capthiron/tima/task"
	"time"
)

type Service interface {
	Tima() *day.Day
	Hi(time time.Time) *day.Day
	Bye(time time.Time) *day.Day
	Task(task *task.Task) *day.Day
}

type DefaultService struct {
	dayService  day.Service
	taskService task.Service
}

func NewDefaultService() *DefaultService {
	return &DefaultService{
		taskService: task.NewService(),
		dayService:  nil,
	}
}
