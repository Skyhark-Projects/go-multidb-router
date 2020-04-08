package lvldb

import (
	"github.com/Skyhark-Projects/go-multidb-router/db"
	"github.com/syndtr/goleveldb/leveldb"
)

type Lvldb struct {
	*leveldb.DB
}

func (l Lvldb) Get(table string, key []byte) (interface{}, error) {
	return nil, nil
}

func (l Lvldb) Put(table string, key []byte, val interface{}) error {
	return nil
}

func (l Lvldb) Delete(table string, key []byte) error {
	return nil
}

// -----------

func Open(uri string) (db.DB, error) {
	lvl, err := leveldb.OpenFile(uri, nil)
	if err != nil {
		return nil, err
	}

	return db.NewKeyValueDb(Lvldb{ lvl }), nil
}

func init() {
	db.Register("lvldb", Open)
}