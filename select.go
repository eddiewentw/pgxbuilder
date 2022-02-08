package pgxbuilder

import (
	"strconv"
	"strings"
)

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

func (q *Query) Limit(v uint64) *Query {
	q.limit = v

	return q
}

func (q *Query) Offset(v uint64) *Query {
	q.offset = v

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

	if q.limit > 0 {
		b.WriteString(" LIMIT ")
		b.WriteString(strconv.FormatUint(q.limit, 10))
	}
	if q.offset > 0 {
		b.WriteString(" OFFSET ")
		b.WriteString(strconv.FormatUint(q.offset, 10))
	}

	return b.String()
}
