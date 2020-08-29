package task

import (
	"fmt"
	"os"
	"time"
)

type Service interface {
	getTasksForDay(day *time.Time) []Task
}

type DefaultService struct {
	repo Repository
}

func NewDefaultService() *DefaultService {
	service := &DefaultService{}

	switch os.Getenv("PROFILE") {
	case "dev":
	case "test":
		fmt.Println("dev/test profile")
		service.repo = &InMemRepository{}
	default:
		service.repo = &DefaultRepository{}
	}

	return service
}

func (s DefaultService) getTasksForDay(day *time.Time) []Task {

	return []Task{}
}
