package pgxbuilder

// GroupBy groups together those rows in a table that have the same values
// in all the columns listed.
func (q *Query) GroupBy(columns ...string) *Query {
	q.groupBy = append(q.groupBy, columns...)

	return q
}
