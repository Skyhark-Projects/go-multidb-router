package mongodb

type Where map[string]interface{}

func (w Where) And(query interface{}, args ...interface{}) Where {
	if _, ok := query.(string); ok {
		// ToDo parse sql statement
	}

	// if map => merge map
	// if struct => merge with non empty fields
	// if list > where primary key in
	// if empty list => skip

	// ToDo
	return w
}

func (w Where) Or(query interface{}, args ...interface{}) Where {
	if _, ok := query.(string); ok {
		// ToDo parse sql statement
	}

	// ToDo
	return w
}

func (w Where) Not(query interface{}, args ...interface{}) Where {
	if _, ok := query.(string); ok {
		// ToDo parse sql statement
	}

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