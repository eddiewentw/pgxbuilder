package pgxbuilder

import "strings"

type clauseDistinct struct {
	enabled bool
	columns []string
}

func (c clauseDistinct) String() string {
	if !c.enabled {
		return ""
	}

	if len(c.columns) == 0 {
		return "DISTINCT "
	}

	var b strings.Builder

	b.WriteString("DISTINCT ON (")
	b.WriteString(strings.Join(c.columns, ", "))
	b.WriteString(") ")

	return b.String()
}

// Distinct excludes all duplicate rows from the result set. One row will be
// kept from each group of duplicates.
func (q *Query) Distinct() *Query {
	q.distinct.enabled = true

	return q
}

// DistinctOn keeps only the first row of each set of rows where the given
// expressions evaluate to equal.
func (q *Query) DistinctOn(columns ...string) *Query {
	q.distinct.enabled = true
	q.distinct.columns = columns

	return q
}
