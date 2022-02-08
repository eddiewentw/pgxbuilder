package pgxbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_GroupBy(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		q := From("members").
			GroupBy("age")

		assert.Equal(t, "SELECT * FROM members GROUP BY age", q.String())
	})

	t.Run("with many columns", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			q := From("members").
				GroupBy("age", "gender")

			assert.Equal(t, "SELECT * FROM members GROUP BY age, gender", q.String())
		})

		t.Run("with many GroupBy() calls", func(t *testing.T) {
			q := From("members").
				GroupBy("age").
				GroupBy("gender")

			assert.Equal(t, "SELECT * FROM members GROUP BY age, gender", q.String())
		})
	})

	t.Run("with conditions", func(t *testing.T) {
		q := From("posts")
		q = q.Where("id = $1", q.Param(299)).
			OrderBy("created_at", "ASC")

		assert.Equal(t, "SELECT * FROM posts WHERE (id = $1) ORDER BY created_at ASC", q.String())
	})
}
