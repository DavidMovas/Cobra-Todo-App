package db

import (
	"cobratodoapp/internal/models"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
	DB    *bolt.DB
	Tasks *models.TaskList
}

func (b *BoltDB) Add(task *models.Task) {

}

func (b *BoltDB) Complete(id uint64) error {
	return nil
}

func (b *BoltDB) Remove(id uint64) error {
	return nil
}
