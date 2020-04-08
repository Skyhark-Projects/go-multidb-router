package db

import (
	"github.com/jinzhu/gorm"
)

type GormDB struct {
	*gorm.DB
}

func (g GormDB) AutoMigrate(values ...interface{}) error {
	db := g.DB.AutoMigrate(values...)
	return db.Error
}

func (g GormDB) Begin() DB {
	return GormDB{
		DB: g.DB.Begin(),
	}
}

func (g GormDB) Commit() error {
	db := g.DB.Commit()
	return db.Error
}

func (g GormDB) Rollback() error {
	db := g.DB.Rollback()
	return db.Error
}

// -----------

func (g GormDB) Count(value interface{}) error {
	db := g.DB.Count(value)
	return db.Error
}

func (g GormDB) Create(value interface{}) error {
	db := g.DB.Create(value)
	return db.Error
}

func (g GormDB) Delete(value interface{}, where ...interface{}) error {
	db := g.DB.Delete(value, where...)
	return db.Error
}

func (g GormDB) Save(value interface{}) error {
	db := g.DB.Save(value)
	return db.Error
}

func (g GormDB) Update(value ...interface{}) error {
	db := g.DB.Update(value...)
	return db.Error
}

// -----------

func (g GormDB) Find(out interface{}, where ...interface{}) error {
	db := g.DB.Find(out, where...)
	return db.Error
}

func (g GormDB) First(out interface{}, where ...interface{}) error {
	db := g.DB.Find(out, where...)
	return db.Error
}

func (g GormDB) FirstOrCreate(out interface{}, where ...interface{}) error {
	db := g.DB.FirstOrCreate(out, where...)
	return db.Error
}

func (g GormDB) Last(out interface{}, where ...interface{}) error {
	db := g.DB.Last(out, where...)
	return db.Error
}

// -----------

func (g GormDB) Preload(column string) DB {
	return GormDB{
		DB: g.DB.Preload(column),
	}
}

func (g GormDB) Model(in interface{}) DB {
	return GormDB{
		DB: g.DB.Model(in),
	}
}

func (g GormDB) Select(query interface{}, args ...interface{}) DB {
	return GormDB{
		DB: g.DB.Select(query, args...),
	}
}

func (g GormDB) Joins(query string, args ...interface{}) DB {
	return GormDB{
		DB: g.DB.Joins(query, args...),
	}
}

func (g GormDB) Group(query string) DB {
	return GormDB{
		DB: g.DB.Group(query),
	}
}

func (g GormDB) Having(query interface{}, values ...interface{}) DB {
	return GormDB{
		DB: g.DB.Having(query, values),
	}
}

// -----------

func (g GormDB) Offset(offset interface{}) DB {
	return GormDB{
		DB: g.DB.Offset(offset),
	}
}

func (g GormDB) Limit(limit interface{}) DB {
	return GormDB{
		DB: g.DB.Limit(limit),
	}
}

func (g GormDB) Table(name string) DB {
	return GormDB{
		DB: g.DB.Table(name),
	}
}

// -----------

func (g GormDB) Not(query interface{}, args ...interface{}) DB {
	return GormDB{
		DB: g.DB.Not(query, args...),
	}
}

func (g GormDB) Or(query interface{}, args ...interface{}) DB {
	return GormDB{
		DB: g.DB.Or(query, args...),
	}
}

func (g GormDB) Where(query interface{}, args ...interface{}) DB {
	return GormDB{
		DB: g.DB.Where(query, args...),
	}
}
