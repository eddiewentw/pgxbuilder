package pgxbuilder

import "strings"

// Delete starts a delete statement.
func Delete(table string) *Query {
	return &Query{
		stmt:  stmtDelete,
		table: table,
	}
}

func (q Query) toDelete() string {
	var b strings.Builder

	b.WriteString("DELETE FROM ")
	b.WriteString(q.table)

	b.WriteString(q.toWhere())

	return b.String()
}
