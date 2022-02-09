package pgxbuilder

import "strings"

type whereCondition struct {
	operator  operator
	condition string
	extra     []whereCondition
}

func (c whereCondition) String() string {
	var b strings.Builder

	b.WriteString(c.operator.String())
	b.WriteString("(")

	b.WriteString(c.condition)

	for _, cond := range c.extra {
		b.WriteString(cond.String())
	}

	b.WriteString(")")

	return b.String()
}

func (c whereCondition) apply(_ *Query, cond *whereCondition) {
	cond.extra = append(cond.extra, c)
}

type whereOption interface {
	// apply adds this option to the main condition.
	apply(*Query, *whereCondition)
}

// Where adds one or more where conditions. Use And and Or operators to combine
// multiple conditions.
func (q *Query) Where(condition string, options ...whereOption) *Query {
	c := whereCondition{
		condition: condition,
	}

	for _, o := range options {
		o.apply(q, &c)
	}

	q.conditions = append(q.conditions, c)

	return q
}

func (q Query) toWhere() string {
	if len(q.conditions) == 0 {
		return ""
	}

	var b strings.Builder

	b.WriteString(" WHERE ")

	for i, c := range q.conditions {
		b.WriteString(c.String())

		if i != len(q.conditions)-1 {
			b.WriteString(operatorAnd.String())
		}
	}

	return b.String()
}
