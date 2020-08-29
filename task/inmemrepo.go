package task

import (
	"math/rand"
)

type InMemRepository struct {
	tasks []Task
}

func (r *InMemRepository) AddTask(task *Task) (*Task, error) {
	task.Model.ID = uint(rand.Uint32())
	r.tasks = append(r.tasks, *task)
	return task, nil
}

func (r InMemRepository) GetTasks() ([]Task, error) {
	return r.tasks, nil
}

func (r *InMemRepository) UpdateTask(taskToUpdate *Task) (*Task, error) {
	for i, task := range r.tasks {
		if task.Model.ID == taskToUpdate.Model.ID {
			r.tasks[i] = *taskToUpdate
		}
	}
	return taskToUpdate, nil
}

func (r *InMemRepository) RemoveTask(id uint) error {
	for i, task := range r.tasks {
		if task.Model.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
		}
	}
	return nil
}
