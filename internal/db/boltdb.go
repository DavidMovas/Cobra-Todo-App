package db

import (
	"cobratodoapp/internal/models"
	"encoding/binary"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
	DB *bolt.DB
}

func NewBoltDB(db *bolt.DB) *BoltDB {
	_ = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		return err
	})

	return &BoltDB{
		DB: db,
	}
}

func (b *BoltDB) Add(task *models.Task) {
	_ = b.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		id, _ := b.NextSequence()
		task.ID = id

		idBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(idBytes, task.ID)
		return b.Put(idBytes, []byte(task.Title))
	})
}

func (b *BoltDB) Complete(id uint64) error {
	return nil
}

func (b *BoltDB) Remove(id uint64) error {
	return nil
}

func (b *BoltDB) readFromSource() (*models.TaskList, error) {

	return nil, nil
}
