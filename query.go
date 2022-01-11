package pgxbuilder

import "strings"

type Query struct {
	// stmt indicates which kind of statement this query is.
	stmt statement

	table string

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
	var b strings.Builder

	switch q.stmt {
	case stmtSelect:
		b.WriteString("SELECT ")

		if len(q.columns) == 0 {
			b.WriteString("*")
		} else {
			b.WriteString(strings.Join(q.columns, ", "))
		}

		b.WriteString(" FROM ")
		b.WriteString(q.table)
	case stmtDelete:
		b.WriteString("DELETE FROM ")
		b.WriteString(q.table)
	}

	return b.String()
}
