package pgxbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_OrderBy(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		q := From("posts")
		q = q.OrderBy("created_at")

		assert.Equal(t, "SELECT * FROM posts ORDER BY created_at", q.String())
	})

	t.Run("with multiple expressions", func(t *testing.T) {
		q := From("posts")
		q = q.OrderBy("created_at", "DESC").
			OrderBy("id")

		assert.Equal(t, "SELECT * FROM posts ORDER BY created_at DESC, id", q.String())
	})

	t.Run("with where clause", func(t *testing.T) {
		q := From("posts")
		q = q.Where("id = $1", q.Param(299)).
			OrderBy("created_at", "ASC")

		assert.Equal(t, "SELECT * FROM posts WHERE (id = $1) ORDER BY created_at ASC", q.String())
	})

	t.Run("with group by clause", func(t *testing.T) {
		q := From("posts").
			GroupBy("member_id").
			OrderBy("created_at", "NULLS FIRST")

		assert.Equal(t, "SELECT * FROM posts GROUP BY member_id ORDER BY created_at NULLS FIRST", q.String())
	})
}
