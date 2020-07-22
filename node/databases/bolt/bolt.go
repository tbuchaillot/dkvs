package bolt

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
)


var defaultBucket = []byte("default")
type Database struct {
	db       *bolt.DB
	closeFunc func() error
}

func NewDatabase(dbPath string) (db *Database, err error) {
	boltDb, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	db = &Database{db: boltDb}
	db.closeFunc = boltDb.Close

	if err := db.createDefaultBucket(); err != nil {
		db.closeFunc()
		return nil, fmt.Errorf("creating default bucket: %w", err)
	}

	return db, nil
}

func (d *Database) createDefaultBucket() error {
	return d.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(defaultBucket)
		return err
	})
}

func (d *Database) SetKey(key string, value []byte) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(defaultBucket)
		return b.Put([]byte(key), value)
	})
}

func (d *Database) GetKey(key string) ([]byte, error) {
	var result []byte
	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(defaultBucket)
		result = b.Get([]byte(key))
		return nil
	})

	if err == nil {
		return result, nil
	}
	return nil, err
}

func (d *Database) Close() error{
	return d.closeFunc()
}
