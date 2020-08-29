package task

type repository interface {
	addTask(task *Task) (*Task, error)
	getTasks() ([]Task, error)
	updateTask(task *Task) (*Task, error)
	removeTask(id uint) error
}

type defaultRepository struct {
}

func (r defaultRepository) addTask(task *Task) (*Task, error) {
	return nil, nil
}

func (r defaultRepository) getTasks() ([]Task, error) {
	return nil, nil
}

func (r defaultRepository) updateTask(task *Task) (*Task, error) {
	return nil, nil
}

func (r defaultRepository) removeTask(id uint) error {
	return nil
}
