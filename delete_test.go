package pgxbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	t.Run("delete all records in the table", func(t *testing.T) {
		q := Delete("posts")

		assert.Equal(t, "DELETE FROM posts", q.String())
	})

	t.Run("with a condition", func(t *testing.T) {
		q := Delete("posts")
		q = q.Where("id = $1", q.Param(299))

		assert.Equal(t, "DELETE FROM posts WHERE (id = $1)", q.String())
		assert.Equal(t, []interface{}{299}, q.Parameters())
	})

	t.Run("return columns", func(t *testing.T) {
		q := Delete("posts")
		q = q.Where("id = $1", q.Param(299)).
			Returning("updated_at", "deleted_at")

		assert.Equal(t, "DELETE FROM posts WHERE (id = $1) RETURNING updated_at, deleted_at", q.String())
	})
}
