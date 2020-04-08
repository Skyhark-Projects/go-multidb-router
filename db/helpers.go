package db

import (
	"github.com/jinzhu/gorm"
	"reflect"
)

type tabler interface {
	TableName() string
}

func TableName(i interface{}) string {
	if tabler, ok := i.(tabler); ok {
		return tabler.TableName()
	} else {
		name := gorm.ToTableName( reflect.TypeOf(i).Elem().Name() )
		if l := len(name); l > 0 && name[l-1] != 's' {
			return name + "s"
		} else {
			return name
		}
	}
}

// --------------

type Where map[string]interface{}

func (w Where) And(query interface{}, args ...interface{}) Where {
	// ToDo
	return w
}

func (w Where) Or(query interface{}, args ...interface{}) Where {
	// ToDo
	return w
}

func (w Where) Not(query interface{}, args ...interface{}) Where {
	// ToDo
	return w
}

func (w Where) Final(where ...interface{}) Where {
	if len(where) > 0 {
		return w.And(where[0], where[1:]...)
	} else {
		return w
	}
}

// --------------

func Int64(val interface{}) int64 {
	if v, ok := val.(int64); ok {
		return v
	} else if v, ok := val.(int32); ok {
		return int64(v)
	} else if v, ok := val.(int16); ok {
		return int64(v)
	} else if v, ok := val.(int8); ok {
		return int64(v)
	} else if v, ok := val.(int); ok {
		return int64(v)
	} else if v, ok := val.(float64); ok {
		return int64(v)
	} else if v, ok := val.(float32); ok {
		return int64(v)
	} else {
		panic("Could not parse int64 type")
	}
}