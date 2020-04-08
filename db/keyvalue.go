package db

import (
	"errors"
)

type KeyValueHandler interface {
	// Has(table string, key []byte) (bool, error)
	Get(table string, key []byte) (interface{}, error)
	Put(table string, key []byte, val interface{}) error
	Delete(table string, key []byte) error
	Close() error
}

type KeyValueDB struct {
	handler KeyValueHandler
}

func NewKeyValueDb(handler KeyValueHandler) KeyValueDB {
	return KeyValueDB{
		handler: handler,
	}
}

func (es KeyValueDB) AutoMigrate(values ...interface{}) error {
	return nil
}

func (es KeyValueDB) Close() error {
	return es.handler.Close()
}

// -----------

func (es KeyValueDB) Begin() DB {
	return nil // ToDo
}

func (es KeyValueDB) Commit() error {
	return errors.New("Not supported yet")
}

func (es KeyValueDB) Rollback() error {
	return errors.New("Not supported yet")
}

// -----------

func (es KeyValueDB) Count(value interface{}) error {
	// ToDo
	return errors.New("not supported yet")
}

func (es KeyValueDB) Create(value interface{}) error {
	// ToDo
	return errors.New("not supported yet")
}

func (es KeyValueDB) Delete(value interface{}, where ...interface{}) error {
	// ToDo
	return errors.New("not supported yet")
}

func (es KeyValueDB) Save(value interface{}) error {
	// ToDo
	return errors.New("not supported yet")
}

func (es KeyValueDB) Update(value ...interface{}) error {
	// ToDo
	return errors.New("not supported yet")
}

// -----------

func (es KeyValueDB) Find(out interface{}, where ...interface{}) error {
	return errors.New("not supported yet")
}

func (es KeyValueDB) First(out interface{}, where ...interface{}) error {
	return errors.New("not supported yet")
}

func (es KeyValueDB) FirstOrCreate(out interface{}, where ...interface{}) error {
	return errors.New("not supported yet")
}

func (es KeyValueDB) Last(out interface{}, where ...interface{}) error {
	return errors.New("not supported yet")
}

// -----------

func (es KeyValueDB) Preload(column string) DB {
	return nil
}

func (es KeyValueDB) Model(in interface{}) DB {
	return nil
}

func (es KeyValueDB) Select(query interface{}, args ...interface{}) DB {
	return nil
}

func (es KeyValueDB) Joins(query string, args ...interface{}) DB {
	return nil
}

func (es KeyValueDB) Group(query string) DB {
	return nil
}

func (es KeyValueDB) Having(query interface{}, values ...interface{}) DB {
	return nil
}

// -----------

func (es KeyValueDB) Offset(offset interface{}) DB {
	return nil
}

func (es KeyValueDB) Limit(limit interface{}) DB {
	return nil
}

func (es KeyValueDB) Table(name string) DB {
	return nil
}

// -----------

func (es KeyValueDB) Not(query interface{}, args ...interface{}) DB {
	return nil
}

func (es KeyValueDB) Or(query interface{}, args ...interface{}) DB {
	return nil
}

func (es KeyValueDB) Where(query interface{}, args ...interface{}) DB {
	return nil
}