package pgxbuilder

import "strings"

// From starts a select statement.
func From(table string) *Query {
	return &Query{
		stmt:  stmtSelect,
		table: table,
	}
}

// Select adds items to the select list.
func (q *Query) Select(columns ...string) *Query {
	q.columns = append(q.columns, columns...)

	return q
}

func (q Query) toSelect() string {
	var b strings.Builder

	b.WriteString("SELECT ")

	if len(q.columns) == 0 {
		b.WriteString("*")
	} else {
		b.WriteString(strings.Join(q.columns, ", "))
	}

	b.WriteString(" FROM ")
	b.WriteString(q.table)

	b.WriteString(q.toWhere())

	return b.String()
}
