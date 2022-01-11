package pgxbuilder

// Delete starts a delete statement.
func Delete(table string) *Query {
	return &Query{
		stmt:  stmtDelete,
		table: table,
	}
}
