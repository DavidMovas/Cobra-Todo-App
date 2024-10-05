package models

import "errors"

var (
	ErrTaskNotFound = errors.New("[ERROR] Task not found")
)

type TaskList []*Task

func (t *TaskList) Add(task *Task) {
	*t = append(*t, task)
}

func (t *TaskList) Remove(id uint64) error {
	for i, task := range *t {
		if task.ID == id {
			*t = append((*t)[:i], (*t)[i+1:]...)
			return nil
		}
	}

	return ErrTaskNotFound
}

func (t *TaskList) Complete(id uint64) error {
	for _, task := range *t {
		if task.ID == id {
			task.Complete()
			return nil
		}
	}

	return ErrTaskNotFound
}
