package memory

import (
	"github.com/Skyhark-Projects/go-multidb-router/db"
)

type Memory struct {
	
}

func (l *Memory) Get(table string, key []byte) (interface{}, error) {
	return nil, nil
}

func (l *Memory) Put(table string, key []byte, val interface{}) error {
	return nil
}

func (l *Memory) Delete(table string, key []byte) error {
	return nil
}

func (l *Memory) Close() error {
	return nil
}

// -----------

func Open(uri string) (db.DB, error) {
	return db.NewKeyValueDb(&Memory{}), nil
}

func init() {
	db.Register("memory", Open)
}