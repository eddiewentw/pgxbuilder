package pgxbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrom(t *testing.T) {
	q := From("posts")

	assert.Equal(t, "SELECT * FROM posts", q.String())
}

func TestQuery_Select(t *testing.T) {
	t.Run("select one column", func(t *testing.T) {
		q := From("posts").
			Select("title")

		assert.Equal(t, "SELECT title FROM posts", q.String())
	})

	t.Run("select many columns", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			q := From("posts").
				Select("title", "content", "author")

			assert.Equal(t, "SELECT title, content, author FROM posts", q.String())
		})

		t.Run("with many Select() calls", func(t *testing.T) {
			q := From("posts").
				Select("title").
				Select("content", "author")

			assert.Equal(t, "SELECT title, content, author FROM posts", q.String())
		})
	})

	t.Run("with a condition", func(t *testing.T) {
		q := From("posts")
		q = q.Where("id = $1", q.Param(299))

		assert.Equal(t, "SELECT * FROM posts WHERE (id = $1)", q.String())
		assert.Equal(t, []interface{}{299}, q.Parameters())
	})
}
