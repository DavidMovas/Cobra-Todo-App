package db

import "cobratodoapp/internal/models"

type DB interface {
	Add(task *models.Task)
	Complete(id uint64) error
	Remove(id uint64) error
}
