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
