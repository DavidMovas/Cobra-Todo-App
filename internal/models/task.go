package models

import "time"

type Task struct {
	ID          uint64    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Done        bool      `json:"done" db:"done"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	CompletedAt time.Time `json:"completed_at" db:"completed_at"`
}

func NewTask(title string) *Task {
	return &Task{
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
}

func (t *Task) Complete() {
	t.Done = true
	t.CompletedAt = time.Now()
}

func (t *Task) GetCompletedAt() string {
	var result string

	if t.CompletedAt.IsZero() {
		result = "-"
	} else {
		result = t.CompletedAt.Format("2006-01-02 15:04:05")
	}

	return result
}
