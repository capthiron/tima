package task

type Repository interface {
	AddTask(task *Task) (*Task, error)
	GetTasks() ([]Task, error)
	UpdateTask(task *Task) (*Task, error)
	RemoveTask(id uint) error
}

type DefaultRepository struct {
}

func (r DefaultRepository) AddTask(task *Task) (*Task, error) {
	return nil, nil
}

func (r DefaultRepository) GetTasks() ([]Task, error) {
	return nil, nil
}

func (r DefaultRepository) UpdateTask(task *Task) (*Task, error) {
	return nil, nil
}

func (r DefaultRepository) RemoveTask(id uint) error {
	return nil
}
