package db_test

import (
	"testing"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "bitbucket.org/skyhark/golang/db/mongodb"
	"bitbucket.org/skyhark/golang/db"
)

func assert(t *testing.T, e error) {
	if e != nil {
		t.Fatal(e)
	}
}

type TestTable struct {
	ID uint
	Name string
}

// ----------------------------------------------------------------------------------------

func execTestes(t *testing.T, db_type string, uri string) {
	db, err := db.Open(db_type, uri)
	assert(t, err)

	// Create test table
	err = db.AutoMigrate(&TestTable{})
	assert(t, err)

	// Delete all rows
	assert(t, db.Delete(&TestTable{}))

	// Create row
	assert(t, db.Create(&TestTable{
		Name: "x",
	}))

	assert(t, db.Save(&TestTable{
		Name: "x2",
	}))
}

// Layers:
// GORM (mysql, postgress, sqllite, mssql)
// Mongodb
// Elasticsearch
// Keyvalue store

// ----------------------------------------------------------------------------------------

func TestMysql(t *testing.T) {
	execTestes(t, "mysql", "root:mysql@tcp(0.0.0.0:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
}

func TestPostgress(t *testing.T) {
	// execTestes(t, "postgress", "")
}

func TestSqlite(t *testing.T) {
	// execTestes(t, "sqlite3", "")
}

func TestMssql(t *testing.T) {
	// execTestes(t, "mssql", "")
}

func TestElasticsearch(t *testing.T) {
	// execTestes(t, "elasticsearch", "")
}

func TestMongodb(t *testing.T) {
	// execTestes(t, "mongodb", "mongodb://localhost:27017/testing")
}

func TestLvldb(t *testing.T) {
	// execTestes(t, "lvldb", ".db")
}

func TestRedis(t *testing.T) {
	// execTestes(t, "redis", "")
}

func TestMemcached(t *testing.T) {
	// execTestes(t, "memcached", "")
}

func TestMemory(t *testing.T) {
	// execTestes(t, "memory", "")
}