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
	b.WriteRune('(')

	b.WriteString(c.condition)

	for _, cond := range c.extra {
		b.WriteString(cond.String())
	}

	b.WriteRune(')')

	return b.String()
}

// Where adds one or more where conditions. Use And and Or operators to combine
// multiple conditions.
func (q *Query) Where(condition string, conditions ...whereCondition) *Query {
	q.conditions = append(q.conditions, whereCondition{
		// operator:  operatorAnd,
		condition: condition,
		extra:     conditions,
	})

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
