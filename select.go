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

// Distinct excludes all duplicate rows from the result set. One row will be
// kept from each group of duplicates.
func (q *Query) Distinct() *Query {
	q.distinct = true

	return q
}

// Limit specifies the maximum number of rows to return.
func (q *Query) Limit(count uint64) *Query {
	q.limit = count

	return q
}

// Offset specifies the number of rows to skip before starting to return rows.
func (q *Query) Offset(start uint64) *Query {
	q.offset = start

	return q
}

func (q Query) toSelect() string {
	var b strings.Builder

	b.WriteString("SELECT ")

	if q.distinct {
		b.WriteString("DISTINCT ")
	}

	if len(q.columns) == 0 {
		b.WriteString("*")
	} else {
		b.WriteString(strings.Join(q.columns, ", "))
	}

	b.WriteString(" FROM ")
	b.WriteString(q.table)

	b.WriteString(q.toWhere())

	if len(q.groupBy) > 0 {
		b.WriteString(" GROUP BY ")
		b.WriteString(strings.Join(q.groupBy, ", "))
	}
	if len(q.orderBy) > 0 {
		b.WriteString(" ORDER BY ")
		b.WriteString(strings.Join(q.orderBy, ", "))
	}
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
