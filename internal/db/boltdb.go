package db

import (
	"cobratodoapp/internal/convert"
	"cobratodoapp/internal/models"
	"encoding/json"

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

func (b *BoltDB) Tasks() *models.TaskList {
	taskList := models.TaskList{}

	_ = b.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		b.ForEach(func(k, v []byte) error {
			task := &models.Task{}
			err := json.Unmarshal(v, task)
			if err != nil {
				return err
			}
			taskList = append(taskList, task)
			return nil
		})
		return nil
	})

	return &taskList
}

func (b *BoltDB) Add(task *models.Task) {
	_ = b.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		id, _ := b.NextSequence()
		task.ID = id

		taskBytes, _ := json.Marshal(task)
		return b.Put(convert.ConvertIDToByte(task.ID), taskBytes)
	})
}

func (b *BoltDB) Complete(id uint64) error {
	return b.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		taskBytes := b.Get(convert.ConvertIDToByte(id))

		task := &models.Task{}
		err := json.Unmarshal(taskBytes, task)
		if err != nil {
			return err
		}

		task.Complete()
		taskBytes, _ = json.Marshal(task)
		return b.Put(convert.ConvertIDToByte(id), taskBytes)
	})
}

func (b *BoltDB) Remove(id uint64) error {
	return b.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("tasks")).Delete(convert.ConvertIDToByte(id))
	})
}
