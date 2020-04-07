package db

import (
	"github.com/jinzhu/gorm"
)

type DB interface {
	AutoMigrate(values ...interface{}) error
	Close() error

	Begin() DB
	Commit() error
	Rollback() error

	Count(value interface{}) error
	Create(value interface{}) error
	Delete(value interface{}, where ...interface{}) error
	Save(value interface{}) error
	Update(attrs ...interface{}) error

	Find(out interface{}, where ...interface{}) error
	First(out interface{}, where ...interface{}) error
	FirstOrCreate(out interface{}, where ...interface{}) error
	Last(out interface{}, where ...interface{}) error

	// Row() *sql.Row
	// Rows() (*sql.Rows, error)
	// GetErrors() []error

	Select(query interface{}, args ...interface{}) DB
	Joins(query string, args ...interface{}) DB
	Group(query string) DB
	Having(query interface{}, values ...interface{}) DB

	Offset(offset interface{}) DB
	Limit(limit interface{}) DB
	Table(name string) DB

	Not(query interface{}, args ...interface{}) DB
	Or(query interface{}, args ...interface{}) DB
	Where(query interface{}, args ...interface{}) DB
}

var types = map[string]func(uri string) (DB, error){}

func Open(db_type string, uri string) (DB, error) {
	if t, ok := types[db_type]; ok {
		return t(uri)
	}

	db, err := gorm.Open(db_type, uri)
	return GormDB{ DB: db }, err
}

func Register(name string, open func(uri string) (DB, error)) {
	types[name] = open
}