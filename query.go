package pgxbuilder

import "strings"

type Query struct {
	// stmt indicates which kind of statement this query is.
	stmt statement

	table      string
	conditions []whereCondition
	parameters []interface{}
	limit      uint64
	offset     uint64
	orderBy    []string
	groupBy    []string
	distinct   clauseDistinct

	// columns has different purposes in different statements. It
	// contains selected columns in select statement, set expressions
	// in update statement.
	columns []string
	// returning contains columns to return after manipulation.
	returning []string

	// valueSize counts how many columns have an explicit value to insert.
	// This helps insert statement to decide how many records are going
	// to insert.
	valueSize int
}

type statement int

const (
	stmtSelect statement = iota + 1
	stmtDelete
	stmtUpdate
	stmtInsert
)

// String returns an SQL string.
func (q Query) String() string {
	switch q.stmt {
	case stmtSelect:
		return q.toSelect()
	case stmtDelete:
		return q.toDelete()
	case stmtUpdate:
		return q.toUpdate()
	case stmtInsert:
		return q.toInsert()
	}

	return ""
}

// Parameters returns all parameters in this query.
func (q Query) Parameters() []interface{} {
	return q.parameters
}

type parameterOption struct{}

func (o parameterOption) apply(_ *Query, _ *whereCondition) {}

// Param registers one or more parameters to this query. You can use Parameters
// to get all of them.
func (q *Query) Param(v ...interface{}) parameterOption {
	q.parameters = append(q.parameters, v...)
	return parameterOption{}
}

// Returning obtains data from modified rows.
func (q *Query) Returning(columns ...string) *Query {
	q.returning = append(q.returning, columns...)

	return q
}

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
