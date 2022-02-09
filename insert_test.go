package pgxbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	t.Run("insert all columns", func(t *testing.T) {
		q := Insert("posts", []string{}).
			Values(1, "foo", "bar")

		assert.Equal(t, "INSERT INTO posts VALUES ($1, $2, $3)", q.String())
		assert.Equal(t, []interface{}{1, "foo", "bar"}, q.Parameters())
	})

	t.Run("insert specific columns", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			q := Insert("posts", []string{"title", "content"}).
				Values("foo", "bar")

			assert.Equal(t, "INSERT INTO posts(title, content) VALUES ($1, $2)", q.String())
			assert.Equal(t, []interface{}{"foo", "bar"}, q.Parameters())
		})

		t.Run("insert multiple records", func(t *testing.T) {
			q := Insert("posts", []string{"title", "content"}).
				Values("foo", "Fly without history, and we won’t teleport a starship.").
				Values("bar", "The planet is bravely crazy.")

			assert.Equal(t, "INSERT INTO posts(title, content) VALUES ($1, $2), ($3, $4)", q.String())
			assert.Equal(t, []interface{}{
				"foo", "Fly without history, and we won’t teleport a starship.",
				"bar", "The planet is bravely crazy.",
			}, q.Parameters())
		})
	})

	t.Run("return columns", func(t *testing.T) {
		q := Insert("posts", []string{}).
			Values(1, "foo", "bar").
			Returning("created_at", "updated_at")

		assert.Equal(t, "INSERT INTO posts VALUES ($1, $2, $3) RETURNING created_at, updated_at", q.String())
	})

	t.Run("with Param() wrapper", func(t *testing.T) {
		q := Insert("posts", []string{"title", "content"})
		q = q.Values(q.Param("foo"), q.Param("Fly without history, and we won’t teleport a starship.")).
			Values(q.Param("bar"), q.Param("The planet is bravely crazy."))

		assert.Equal(t, "INSERT INTO posts(title, content) VALUES ($1, $2), ($3, $4)", q.String())
		assert.Equal(t, []interface{}{
			"foo", "Fly without history, and we won’t teleport a starship.",
			"bar", "The planet is bravely crazy.",
		}, q.Parameters())
	})
}
