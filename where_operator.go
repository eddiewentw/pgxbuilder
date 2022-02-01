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

func (q *Query) And(condition string, conditions ...whereCondition) whereCondition {
	return whereCondition{
		operator:  operatorAnd,
		condition: condition,
		extra:     conditions,
	}
}

func (q *Query) Or(condition string, conditions ...whereCondition) whereCondition {
	return whereCondition{
		operator:  operatorOr,
		condition: condition,
		extra:     conditions,
	}
}
