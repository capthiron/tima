package task

import (
	"os"
	"time"
)

type Service interface {
	getTasksForDay(day *time.Time) []Task
}

func NewService() Service {
	service := &defaultService{}

	switch os.Getenv("PROFILE") {
	case "dev":
	case "test":
		service.repo = &inMemRepository{}
	default:
		service.repo = &defaultRepository{}
	}

	return service
}

type defaultService struct {
	repo repository
}

func (s defaultService) getTasksForDay(day *time.Time) []Task {
	return []Task{}
}
