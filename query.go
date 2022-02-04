package pgxbuilder

type Query struct {
	// stmt indicates which kind of statement this query is.
	stmt statement

	table      string
	conditions []whereCondition
	parameters []interface{}

	// columns has different purposes in different statements. It
	// contains selected columns in select statement, set expressions
	// in update statement.
	columns []string
}

type statement int

const (
	stmtSelect statement = iota + 1
	stmtDelete
	stmtUpdate
)

// String returns an SQL string.
func (q Query) String() string {
	switch q.stmt {
	case stmtSelect:
		return q.toSelect()
	case stmtDelete:
		return q.toDelete()
	case stmtUpdate:
		return q.toUpdate()
	}

	return ""
}

// Parameters returns all parameters in this query.
func (q Query) Parameters() []interface{} {
	return q.parameters
}
