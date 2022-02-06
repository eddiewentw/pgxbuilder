package pgxbuilder

import "strings"

// Update starts a update statement.
func Update(table string) *Query {
	return &Query{
		stmt:  stmtUpdate,
		table: table,
	}
}

// Set adds a set clause.
func (q *Query) Set(set string, params ...interface{}) *Query {
	q.columns = append(q.columns, set)
	q.parameters = append(q.parameters, params...)

	return q
}

func (q Query) toUpdate() string {
	var b strings.Builder

	b.WriteString("UPDATE ")
	b.WriteString(q.table)

	b.WriteString(" SET ")

	for i, s := range q.columns {
		b.WriteString(s)

		if i != len(q.columns)-1 {
			b.WriteString(", ")
		}
	}

	b.WriteString(q.toWhere())

	if len(q.returning) > 0 {
		b.WriteString(" RETURNING ")
		b.WriteString(strings.Join(q.returning, ", "))
	}

	return b.String()
}
