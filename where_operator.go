package pgxbuilder

type operator int

const (
	operatorAnd operator = iota + 1
	operatorOr
)

func (o operator) String() string {
	if o == operatorAnd {
		return " AND "
	}

	if o == operatorOr {
		return " OR "
	}

	return ""
}

func (q *Query) newWhereConditionOption(op operator, condition string, options ...whereOption) whereCondition {
	c := whereCondition{
		operator:  op,
		condition: condition,
	}

	for _, o := range options {
		o.apply(q, &c)
	}

	return c
}

// And concatenates where conditions with an AND operator.
func (q *Query) And(condition string, options ...whereOption) whereCondition {
	return q.newWhereConditionOption(operatorAnd, condition, options...)
}

// Or concatenates where conditions with an OR operator.
func (q *Query) Or(condition string, options ...whereOption) whereCondition {
	return q.newWhereConditionOption(operatorOr, condition, options...)
}
