package memcached

import (
	"github.com/Skyhark-Projects/go-multidb-router/db"
	"github.com/bradfitz/gomemcache/memcache"
	"strings"
)

type Memcached struct {
	*memcache.Client
}

func (l Memcached) Get(table string, key []byte) (interface{}, error) {
	return nil, nil
}

func (l Memcached) Put(table string, key []byte, val interface{}) error {
	return nil
}

func (l Memcached) Delete(table string, key []byte) error {
	return nil
}

func (l Memcached) Close() error {
	return nil
}

// -----------

func Open(uri string) (db.DB, error) {
	mc := memcache.New(strings.Split(uri, ",")...)
	return db.NewKeyValueDb(Memcached{ mc }), nil
}

func init() {
	db.Register("memcached", Open)
}