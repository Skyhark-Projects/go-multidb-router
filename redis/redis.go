package redis

import (
	"github.com/Skyhark-Projects/go-multidb-router/db"
	// "github.com/go-redis/redis/v7"
)

type Redis struct {

}

func (l Redis) Get(table string, key []byte) (interface{}, error) {
	return nil, nil
}

func (l Redis) Put(table string, key []byte, val interface{}) error {
	return nil
}

func (l Redis) Delete(table string, key []byte) error {
	return nil
}

func (l Redis) Close() error {
	return nil
}

// -----------

func Open(uri string) (db.DB, error) {
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	return db.NewKeyValueDb(Redis{ }), nil
}

func init() {
	db.Register("redis", Open)
}