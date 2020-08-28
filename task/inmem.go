package task

import (
	"github.com/capthiron/tima/model"

	"github.com/rs/xid"
)

type InMemRepository struct {
	tasks []model.Task
}

func(s *InMemRepository) AddTask(task *model.Task) (*model.Task, error) {
	task.Id = xid.New().String()
	s.tasks = append(s.tasks, *task)
	return task, nil
}

func(s InMemRepository) GetTasks() ([]model.Task, error) {
	return s.tasks, nil
}

func(s *InMemRepository) RemoveTask(id string) error {
	for i, task := range s.tasks {
		if task.Id == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
		}
	}
	return nil
}
