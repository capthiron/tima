package task

import (
	"testing"
	"time"
)

var inMemRepo *inMemRepository

func setup() {
	inMemRepo = &inMemRepository{}
}

func TestInMemNewTask(t *testing.T) {
	setup()

	name := "Test task"
	start := time.Now()

	task := &Task{}
	task.Name = name
	task.StartTime = start

	newTask, err := inMemRepo.addTask(task)
	if err != nil {
		t.Errorf("Received the following error from inMemRepo.addTask():\n%v", err)
	}

	if newTask.Model.ID == 0 {
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

	tasks, err := inMemRepo.getTasks()
	if err != nil {
		t.Errorf("Received the following error from inMemRepo.getTasks():\n%v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("service.getTasks() should have length of 0 before adding a task.")
	}

	_, _ = inMemRepo.addTask(&Task{})

	tasks, _ = inMemRepo.getTasks()

	if len(tasks) != 1 {
		t.Errorf("service.getTasks() should have length of 1 after adding a task, but got %v.", len(tasks))
	}
}

func TestInMemUpdateTask(t *testing.T) {
	setup()

	originalName := "Axel Stall"
	originalNote := "Best feeling!!"

	originalTask := &Task{Name: originalName, Note: originalNote}

	originalTask, _ = inMemRepo.addTask(originalTask)

	updatedName := "50-50 Grind"
	updatedNote := "Crazy feeling!!!"

	originalTask.Name = updatedName
	originalTask.Note = updatedNote

	updatedTask, err := inMemRepo.updateTask(originalTask)
	if err != nil {
		t.Errorf("Received the following error from inMemRepo.updateTask():\n%v", err)
	}

	if updatedTask.Name != updatedName {
		t.Errorf("updatedTask.Name should be %v, but got %v", updatedName, updatedTask.Name)
	}

	if updatedTask.Note != updatedNote {
		t.Errorf("updatedTask.Note should be %v, but got %v", updatedNote, updatedTask.Note)
	}
}

func TestInMemRemoveTask(t *testing.T) {
	setup()

	newTask, _ := inMemRepo.addTask(&Task{})

	_ = inMemRepo.removeTask(newTask.Model.ID)

	tasks, _ := inMemRepo.getTasks()

	if len(tasks) != 0 {
		t.Errorf("inMemRepo.getTasks() should have length of 0 after removing the only task, but got %v.", len(tasks))
	}
}
