package pgxbuilder

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
