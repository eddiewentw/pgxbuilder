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
		q := Delete("posts").
			Where("id = 299")

		assert.Equal(t, "DELETE FROM posts WHERE (id = 299)", q.String())
	})
}
