package pgxbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		q := Update("posts").
			Set("content = 'Hello, world.'")

		assert.Equal(t, "UPDATE posts SET content = 'Hello, world.'", q.String())
		assert.Empty(t, q.Parameters())
	})

	t.Run("update multiple columns", func(t *testing.T) {
		q := Update("posts").
			Set("title = 'foo'").
			Set("content = 'Hello, world.'")

		assert.Equal(t, "UPDATE posts SET title = 'foo', content = 'Hello, world.'", q.String())
		assert.Empty(t, q.Parameters())
	})

	t.Run("with parameters", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			q := Update("posts").
				Set("content = $1", "Hello, world.")

			assert.Equal(t, "UPDATE posts SET content = $1", q.String())
			assert.Equal(t, []interface{}{"Hello, world."}, q.Parameters())
		})

		t.Run("multiple parameters", func(t *testing.T) {
			t.Run("in one call", func(t *testing.T) {
				q := Update("posts").
					Set("title = $1, content = $2", "foo", "Hello, world.")

				assert.Equal(t, "UPDATE posts SET title = $1, content = $2", q.String())
				assert.Equal(t, []interface{}{"foo", "Hello, world."}, q.Parameters())
			})

			t.Run("in multiple calls", func(t *testing.T) {
				q := Update("posts").
					Set("title = $1", "foo").
					Set("content = $2", "Hello, world.")

				assert.Equal(t, "UPDATE posts SET title = $1, content = $2", q.String())
				assert.Equal(t, []interface{}{"foo", "Hello, world."}, q.Parameters())
			})
		})
	})

	t.Run("with where clause", func(t *testing.T) {
		q := Update("posts").
			Set("content = $1", "Hello, world.")
		q = q.Where("id = $2", q.Param(299))

		assert.Equal(t, "UPDATE posts SET content = $1 WHERE (id = $2)", q.String())
		assert.Equal(t, []interface{}{"Hello, world.", 299}, q.Parameters())
	})

	t.Run("return columns", func(t *testing.T) {
		q := Update("posts").
			Set("content = $1", "Hello, world.").
			Returning("created_at", "updated_at")

		assert.Equal(t, "UPDATE posts SET content = $1 RETURNING created_at, updated_at", q.String())
	})
}
