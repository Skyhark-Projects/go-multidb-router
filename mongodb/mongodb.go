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

	where db.Where
	table string
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
	return nil // ToDo
}

func (m MongoDB) Commit() error {
	return errors.New("Not supported yet")
}

func (m MongoDB) Rollback() error {
	return errors.New("Not supported yet")
}

// -----------

func (m MongoDB) Count(output interface{}) error {
	collection := m.Collection( m.table )
	count, err := collection.CountDocuments(context.Background(), m.filters())

	if err == nil {
		reflect.ValueOf(output).Elem().SetInt(count)
	}

	return err
}

func (m MongoDB) Create(value interface{}) error {
	if value == nil {
		return nil
	}

	collection := m.Collection( db.TableName(value) )
	_, err := collection.InsertOne(context.Background(), value)
	// id := res.InsertedID

	// ToDo support primary key field & auto increment
	// ToDo parse gorm tags to apply to mongo creation

	return err
}

func (m MongoDB) Delete(value interface{}, where ...interface{}) error {
	collection := m.Collection( db.TableName(value) )
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
	collection := m.Collection( db.TableName(out) )

	// ToDo merge where clause
	cursor, err := collection.Find(context.Background(), m.filters())
	if err != nil {
		return err
	}

	return cursor.All(context.Background(), out)
}

func (m MongoDB) First(out interface{}, where ...interface{}) error {
	collection := m.Collection( db.TableName(out) )

	// ToDo merge where clause
	return collection.FindOne(context.Background(), m.filters()).Decode(out)
}

func (m MongoDB) FirstOrCreate(out interface{}, where ...interface{}) error {
	return errors.New("not supported yet")
}

func (m MongoDB) Last(out interface{}, where ...interface{}) error {
	return errors.New("not supported yet")
}

// -----------

func (m MongoDB) Preload(column string) db.DB {
	return nil
}

func (m MongoDB) Model(in interface{}) db.DB {
	return nil
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
	return nil
}

func (m MongoDB) Limit(limit interface{}) db.DB {
	return nil
}

func (m MongoDB) Table(name string) db.DB {
	db := m.copy()
	db.table = name
	return db
}

// -----------

func (m MongoDB) Not(query interface{}, args ...interface{}) db.DB {
	db := m.copy()
	db.where = m.where.Not(query, args...)
	return db
}

func (m MongoDB) Or(query interface{}, args ...interface{}) db.DB {
	db := m.copy()
	db.where = m.where.Or(query, args...)
	return db
}

func (m MongoDB) Where(query interface{}, args ...interface{}) db.DB {
	db := m.copy()
	db.where = m.where.And(query, args...)
	return db
}

func (m MongoDB) copy() MongoDB {
	return MongoDB{
		Database: m.Database,
		where: 	  m.where,
		table: 	  m.table,
	}
}

func (m MongoDB) filters() interface{} {
	return /*m.where*/ map[string]string{} // ToDo convert where clause to mongodb filter
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