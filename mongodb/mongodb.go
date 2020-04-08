package mongodb

import (
	"github.com/Skyhark-Projects/go-multidb-router/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"strings"
	"errors"
	"reflect"
)

type MongoDB struct {
	*mongo.Database

	where Where
	table *mongo.Collection
	options options.FindOptions
}

// -----------

func (m MongoDB) AutoMigrate(values ...interface{}) error {
	// ToDo
	return nil
}

func (m MongoDB) Close() error {
	return m.Client().Disconnect(context.Background())
}

// -----------

func (m MongoDB) Begin() db.DB {
	return m // ToDo prepare bulk query
}

func (m MongoDB) Commit() error {
	return errors.New("Not supported yet")
}

func (m MongoDB) Rollback() error {
	return errors.New("Not supported yet")
}

// -----------

func (m MongoDB) Count(output interface{}) error {
	count, err := m.table.CountDocuments(context.Background(), m.where)
	if err == nil {
		reflect.ValueOf(output).Elem().SetInt(count)
	}

	return err
}

func (m MongoDB) Create(value interface{}) error {
	if value == nil {
		return nil
	}

	collection := m.table
	if collection == nil {
		collection = m.Collection( db.TableName(value) )
	}

	_, err := collection.InsertOne(context.Background(), value)
	// id := res.InsertedID

	// ToDo support primary key field & auto increment
	// ToDo parse gorm tags to apply to mongo creation

	return err
}

func (m MongoDB) Delete(value interface{}, where ...interface{}) error {
	collection := m.table
	if collection == nil {
		collection = m.Collection( db.TableName(value) )
	}

	_, err := collection.DeleteMany(context.Background(), m.where.Final(where...), nil)
	return err
}

func (m MongoDB) Save(value interface{}) error {
	// Update or create
	// ToDo
	return errors.New("not supported yet")
}

func (m MongoDB) Update(value ...interface{}) error {
	// ToDo
	return errors.New("not supported yet")
}

// -----------

func (m MongoDB) Find(out interface{}, where ...interface{}) error {
	collection := m.table
	if collection == nil {
		collection = m.Collection( db.TableName(out) )
	}

	cursor, err := collection.Find(context.Background(), m.where.Final(where...), &m.options)
	if err != nil {
		return err
	}

	return cursor.All(context.Background(), out)
}

func (m MongoDB) First(out interface{}, where ...interface{}) error {
	collection := m.table
	if collection == nil {
		collection = m.Collection( db.TableName(out) )
	}

	op := &options.FindOneOptions{
		Skip: m.options.Skip,
		Sort: m.options.Sort,
		Projection: m.options.Projection,
	}

	return collection.FindOne(context.Background(), m.where.Final(where...), op).Decode(out)
}

func (m MongoDB) FirstOrCreate(out interface{}, where ...interface{}) error {
	if err := m.First(out); err != nil && err.Error() == "ErrNoDocuments" {
		return m.Create(out)
	} else if err != nil {
		return err
	}

	return nil
}

func (m MongoDB) Last(out interface{}, where ...interface{}) error {
	m.options.Sort = map[string]int{ "$natural": -1 }
	return m.First(out, where)
}

// -----------

func (m MongoDB) Preload(column string) db.DB {
	return nil
}

func (m MongoDB) Model(in interface{}) db.DB {
	m.table = m.Collection( db.TableName(in) )
	return m
}

func (m MongoDB) Select(query interface{}, args ...interface{}) db.DB {
	return nil
}

func (m MongoDB) Joins(query string, args ...interface{}) db.DB {
	return nil
}

func (m MongoDB) Group(query string) db.DB {
	return nil
}

func (m MongoDB) Having(query interface{}, values ...interface{}) db.DB {
	return nil
}

// -----------

func (m MongoDB) Offset(offset interface{}) db.DB {
	of := db.Int64(offset)
	m.options.Skip = &of
	return m
}

func (m MongoDB) Limit(limit interface{}) db.DB {
	of := db.Int64(limit)
	m.options.Limit = &of
	return m
}

func (m MongoDB) Table(name string) db.DB {
	m.table = m.Collection(name)
	return m
}

// -----------

func (m MongoDB) Not(query interface{}, args ...interface{}) db.DB {
	m.where = m.where.Not(query, args...)
	return m
}

func (m MongoDB) Or(query interface{}, args ...interface{}) db.DB {
	m.where = m.where.Or(query, args...)
	return m
}

func (m MongoDB) Where(query interface{}, args ...interface{}) db.DB {
	m.where = m.where.And(query, args...)
	return m
}

// -----------

func Open(uri string) (db.DB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.Background()
	err = client.Connect(ctx)

	db := uri
	if i := strings.Index(db, "@"); i != -1 {
		db = db[i+1:]
	} else if i := strings.Index(db, "://"); i != -1 {
		db = db[i+3:]
	}

	if i := strings.Index(db, "/"); i != -1 {
		db = db[i+1:]
	} else {
		return nil, errors.New("No database selected")
	}

	if i := strings.Index(db, "?"); i != -1 {
		db = db[:i]
	}

	return MongoDB{
		Database: client.Database(db),
	}, nil
}

func init() {
	db.Register("mongodb", Open)
}