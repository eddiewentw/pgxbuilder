package pgxbuilder

import "strings"

// OrderBy add an expression to sort result rows.
func (q *Query) OrderBy(expression string, options ...string) *Query {
	if len(options) == 0 {
		q.orderBy = append(q.orderBy, expression)

		return q
	}

	var b strings.Builder
	b.WriteString(expression)

	for _, o := range options {
		b.WriteString(" ")
		b.WriteString(o)
	}

	q.orderBy = append(q.orderBy, b.String())

	return q
}
