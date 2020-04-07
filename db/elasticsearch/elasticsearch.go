package mongodb

import (
	"bitbucket.org/skyhark/golang/db"
	"errors"
)

type Elasticsearch struct {
	uri   string
	where db.Where
}

func (es Elasticsearch) AutoMigrate(values ...interface{}) error {
	// ToDo
	return errors.New("not supported yet")
}

func (es Elasticsearch) Close() error {
	return nil
}

// -----------

func (es Elasticsearch) Begin() db.DB {
	return nil // ToDo
}

func (es Elasticsearch) Commit() error {
	return errors.New("Not supported yet")
}

func (es Elasticsearch) Rollback() error {
	return errors.New("Not supported yet")
}

// -----------

func (es Elasticsearch) Count(value interface{}) error {
	// ToDo
	return errors.New("not supported yet")
}

func (es Elasticsearch) Create(value interface{}) error {
	// ToDo
	return errors.New("not supported yet")
}

func (es Elasticsearch) Delete(value interface{}, where ...interface{}) error {
	// ToDo
	return errors.New("not supported yet")
}

func (es Elasticsearch) Save(value interface{}) error {
	// ToDo
	return errors.New("not supported yet")
}

func (es Elasticsearch) Update(value ...interface{}) error {
	// ToDo
	return errors.New("not supported yet")
}

// -----------

func (es Elasticsearch) Find(out interface{}, where ...interface{}) error {
	return errors.New("not supported yet")
}

func (es Elasticsearch) First(out interface{}, where ...interface{}) error {
	return errors.New("not supported yet")
}

func (es Elasticsearch) FirstOrCreate(out interface{}, where ...interface{}) error {
	return errors.New("not supported yet")
}

func (es Elasticsearch) Last(out interface{}, where ...interface{}) error {
	return errors.New("not supported yet")
}

// -----------

func (es Elasticsearch) Select(query interface{}, args ...interface{}) db.DB {
	return nil
}

func (es Elasticsearch) Joins(query string, args ...interface{}) db.DB {
	return nil
}

func (es Elasticsearch) Group(query string) db.DB {
	return nil
}

func (es Elasticsearch) Having(query interface{}, values ...interface{}) db.DB {
	return nil
}

// -----------

func (es Elasticsearch) Offset(offset interface{}) db.DB {
	return nil
}

func (es Elasticsearch) Limit(limit interface{}) db.DB {
	return nil
}

func (es Elasticsearch) Table(name string) db.DB {
	return nil
}

// -----------

func (es Elasticsearch) Not(query interface{}, args ...interface{}) db.DB {
	return Elasticsearch{
		uri:   es.uri,
		where: es.where.Not(query, args...),
	}
}

func (es Elasticsearch) Or(query interface{}, args ...interface{}) db.DB {
	return Elasticsearch{
		uri:   es.uri,
		where: es.where.Or(query, args...),
	}
}

func (es Elasticsearch) Where(query interface{}, args ...interface{}) db.DB {
	return Elasticsearch{
		uri:   es.uri,
		where: es.where.And(query, args...),
	}
}

// -----------

func Open(uri string) (db.DB, error) {
	return Elasticsearch{
		uri: uri,
	}, nil
}

func init() {
	db.Register("elasticsearch", Open)
}