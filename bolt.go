package bluecache

import (
	"errors"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

type boltBackend struct {
	db *bolt.DB
}

func (m *boltBackend) Init() {
	db, err := bolt.Open("./boltcache.db", 744, nil)
	if err != nil {
		fmt.Println(err)
	}
	m.db = db
}

func (m *boltBackend) Set(k, v []byte) error {
	err := m.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("cache"))
		if err != nil {
			return err
		}
		return b.Put(k, v)
	})

	return err
}

func (m *boltBackend) SetEx(k, v []byte, ex time.Duration) error {
	return nil
}

func (m *boltBackend) Get(k []byte) (v []byte, err error) {
	var bts []byte
	m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("cache"))
		if b == nil {
			return errors.New("Error getting bucket")
		}
		bts = b.Get(k)
		return nil
	})

	if len(bts) == 0 {
		return nil, errors.New("Not Found")
	}
	return bts, nil
}

func (m *boltBackend) Del(k []byte) error {
	err := m.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("cache"))
		if err != nil {
			return err
		}
		return b.Delete(k)
	})
	return err
}

func (m *boltBackend) Clear() error {

	err := m.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte("cache"))
	})
	return err
}

func (m *boltBackend) Close() error {
	if m.db != nil {
		return m.db.Close()
	}
	return nil
}
