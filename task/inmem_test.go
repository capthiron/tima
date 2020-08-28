package task

import (
	"testing"
	"time"

	"github.com/capthiron/tima/model"
)

var inMemRepo *InMemRepository

func setup() {
	inMemRepo = &InMemRepository{}
}

func TestInMemNewTask(t *testing.T) {
	setup()

	name := "Test task"
	start := time.Now()
	
	task := &model.Task{}
	task.Name = name
	task.StartTime = start

	newTask, err := inMemRepo.AddTask(task)
	if err != nil {
		t.Errorf("Received the following error from service.NewTask():\n%v", err)
	}

	if newTask.Id == "" {
		t.Errorf("Task id has no value.")
	}

	if newTask.Name != name {
		t.Errorf("Name should be %v, but is %v instead.", name, newTask.Name)
	}

	if newTask.StartTime != start {
		t.Errorf("Start time does not match.")
	}
}

func TestInMemGetTasks(t *testing.T) {
	setup()

	tasks, err := inMemRepo.GetTasks()
	if err != nil {
		t.Errorf("Received the following error from service.getTasks():\n%v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("service.GetTasks() should have length of 0 before adding a task.")
	}

	inMemRepo.AddTask(&model.Task{})

	tasks, _ = inMemRepo.GetTasks()

	if len(tasks) != 1 {
		t.Errorf("service.GetTasks() should have length of 1 after adding a task, but got %v.", len(tasks))
	}
}

func TestInMemRemoveTask(t *testing.T) {
	setup()

	newTask, _ := inMemRepo.AddTask(&model.Task{})

	_ = inMemRepo.RemoveTask(newTask.Id)

	tasks, _ := inMemRepo.GetTasks()

	if len(tasks) != 0 {
		t.Errorf("service.GetTasks() should have length of 0 after removing the only task, but got %v.", len(tasks))
	}
}
