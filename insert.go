package pgxbuilder

import (
	"strconv"
	"strings"
)

// Insert starts a insert statement.
func Insert(table string, columns []string) *Query {
	return &Query{
		stmt:    stmtInsert,
		table:   table,
		columns: columns,
	}
}

// Values registers values of new records.
func (q *Query) Values(params ...interface{}) *Query {
	for _, p := range params {
		if _, ok := p.(whereOption); !ok {
			q.parameters = append(q.parameters, p)
		}
	}

	q.valueSize = len(params)

	return q
}

func (q Query) toInsert() string {
	var b strings.Builder

	b.WriteString("INSERT INTO ")
	b.WriteString(q.table)

	if len(q.columns) > 0 {
		b.WriteString("(")
		b.WriteString(strings.Join(q.columns, ", "))
		b.WriteString(")")
	}

	b.WriteString(" VALUES ")

	if len(q.parameters) > 0 {
		count := len(q.parameters) / q.valueSize

		for i := 0; i < count; i++ {
			b.WriteString("(")

			for j := 1; j <= q.valueSize; j++ {
				b.WriteString("$")
				b.WriteString(strconv.Itoa(i*count + j))

				if j != q.valueSize {
					b.WriteString(", ")
				}
			}

			b.WriteString(")")

			if i != count-1 {
				b.WriteString(", ")
			}
		}
	}

	if len(q.returning) > 0 {
		b.WriteString(" RETURNING ")
		b.WriteString(strings.Join(q.returning, ", "))
	}

	return b.String()
}
