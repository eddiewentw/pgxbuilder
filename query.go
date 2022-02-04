package pgxbuilder

type Query struct {
	// stmt indicates which kind of statement this query is.
	stmt statement

	table      string
	conditions []whereCondition

	// columns contains selected columns in select statement.
	columns []string
}

type statement int

const (
	stmtSelect statement = iota + 1
	stmtDelete
)

// String returns an SQL string.
func (q Query) String() string {
	switch q.stmt {
	case stmtSelect:
		return q.toSelect()
	case stmtDelete:
		return q.toDelete()
	}

	return ""
}
