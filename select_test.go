package pgxbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrom(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		q := From("posts")

		assert.Equal(t, "SELECT * FROM posts", q.String())
	})

	t.Run("with a condition", func(t *testing.T) {
		q := From("posts")
		q = q.Where("id = $1", q.Param(299))

		assert.Equal(t, "SELECT * FROM posts WHERE (id = $1)", q.String())
		assert.Equal(t, []interface{}{299}, q.Parameters())
	})
}

func TestQuery_Select(t *testing.T) {
	t.Run("select one column", func(t *testing.T) {
		q := From("posts").
			Select("title")

		assert.Equal(t, "SELECT title FROM posts", q.String())
	})

	t.Run("select multiple columns", func(t *testing.T) {
		t.Run("in one call", func(t *testing.T) {
			q := From("posts").
				Select("title", "content", "author")

			assert.Equal(t, "SELECT title, content, author FROM posts", q.String())
		})

		t.Run("in multiple calls", func(t *testing.T) {
			q := From("posts").
				Select("title").
				Select("content", "author")

			assert.Equal(t, "SELECT title, content, author FROM posts", q.String())
		})
	})
}

func TestQuery_Limit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		q := From("posts").
			Limit(30)

		assert.Equal(t, "SELECT * FROM posts LIMIT 30", q.String())
	})

	t.Run("with where clause", func(t *testing.T) {
		q := From("posts")
		q = q.Where("member_id = $1", q.Param(55)).
			Limit(30)

		assert.Equal(t, "SELECT * FROM posts WHERE (member_id = $1) LIMIT 30", q.String())
	})

	t.Run("with group by clause", func(t *testing.T) {
		q := From("posts").
			GroupBy("member_id").
			Limit(30)

		assert.Equal(t, "SELECT * FROM posts GROUP BY member_id LIMIT 30", q.String())
	})

	t.Run("with order by clause", func(t *testing.T) {
		q := From("posts").
			OrderBy("created_at").
			Limit(30)

		assert.Equal(t, "SELECT * FROM posts ORDER BY created_at LIMIT 30", q.String())
	})
}

func TestQuery_Offset(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		q := From("posts").
			Offset(10)

		assert.Equal(t, "SELECT * FROM posts OFFSET 10", q.String())
	})

	t.Run("with where clause", func(t *testing.T) {
		q := From("posts")
		q = q.Where("member_id = $1", q.Param(55)).
			Offset(10)

		assert.Equal(t, "SELECT * FROM posts WHERE (member_id = $1) OFFSET 10", q.String())
	})

	t.Run("with group by clause", func(t *testing.T) {
		q := From("posts").
			GroupBy("member_id").
			Offset(10)

		assert.Equal(t, "SELECT * FROM posts GROUP BY member_id OFFSET 10", q.String())
	})

	t.Run("with order by clause", func(t *testing.T) {
		q := From("posts").
			OrderBy("created_at").
			Offset(10)

		assert.Equal(t, "SELECT * FROM posts ORDER BY created_at OFFSET 10", q.String())
	})

	t.Run("with limit clause", func(t *testing.T) {
		q := From("posts").
			Limit(30).
			Offset(10)

		assert.Equal(t, "SELECT * FROM posts LIMIT 30 OFFSET 10", q.String())
	})
}

func TestQuery_For(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		q := From("posts").
			For("NO KEY UPDATE")

		assert.Equal(t, "SELECT * FROM posts FOR NO KEY UPDATE", q.String())
	})

	t.Run("with where clause", func(t *testing.T) {
		q := From("posts")
		q = q.Where("member_id = $1", q.Param(55)).
			For("KEY SHARE")

		assert.Equal(t, "SELECT * FROM posts WHERE (member_id = $1) FOR KEY SHARE", q.String())
	})

	t.Run("with group by clause", func(t *testing.T) {
		q := From("posts").
			GroupBy("member_id").
			For("NO KEY UPDATE")

		assert.Equal(t, "SELECT * FROM posts GROUP BY member_id FOR NO KEY UPDATE", q.String())
	})

	t.Run("with order by clause", func(t *testing.T) {
		q := From("posts").
			OrderBy("created_at").
			For("SHARE")

		assert.Equal(t, "SELECT * FROM posts ORDER BY created_at FOR SHARE", q.String())
	})

	t.Run("with limit clause", func(t *testing.T) {
		q := From("posts").
			Limit(30).
			For("UPDATE")

		assert.Equal(t, "SELECT * FROM posts LIMIT 30 FOR UPDATE", q.String())
	})

	t.Run("with offset clause", func(t *testing.T) {
		q := From("posts").
			Offset(10).
			For("KEY SHARE")

		assert.Equal(t, "SELECT * FROM posts OFFSET 10 FOR KEY SHARE", q.String())
	})
}
