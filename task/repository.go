package task

import (
	"github.com/capthiron/tima/model"
)

type Repository interface {
	AddTask(task *model.Task) (*model.Task, error)
	GetTasks() ([]model.Task, error)
	RemoveTask(id string) error
}

