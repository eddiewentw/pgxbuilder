package pgxbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhere(t *testing.T) {
	t.Run("with one condition", func(t *testing.T) {
		t.Run("without parameters", func(t *testing.T) {
			q := From("posts").
				Where("id = 300")

			assert.Equal(t, "SELECT * FROM posts WHERE (id = 300)", q.String())
			assert.Empty(t, q.Parameters())
		})

		t.Run("with one parameter", func(t *testing.T) {
			q := From("posts")
			q = q.Where("id = $1", q.Param(300))

			assert.Equal(t, "SELECT * FROM posts WHERE (id = $1)", q.String())
			assert.Equal(t, []interface{}{300}, q.Parameters())
		})

		t.Run("with multiple parameters", func(t *testing.T) {
			t.Run("in one call", func(t *testing.T) {
				q := From("posts")
				q = q.Where("id = $1 OR id = $2", q.Param(300, 301))

				assert.Equal(t, "SELECT * FROM posts WHERE (id = $1 OR id = $2)", q.String())
				assert.Equal(t, []interface{}{300, 301}, q.Parameters())
			})

			t.Run("in multiple calls", func(t *testing.T) {
				q := From("posts")
				q = q.Where("id = $1 OR id = $2", q.Param(300), q.Param(301))

				assert.Equal(t, "SELECT * FROM posts WHERE (id = $1 OR id = $2)", q.String())
				assert.Equal(t, []interface{}{300, 301}, q.Parameters())
			})
		})
	})

	t.Run("with multiple conditions", func(t *testing.T) {
		t.Run("in one call", func(t *testing.T) {
			q := From("members")
			q = q.Where("id = $1 AND deleted_at IS NULL", q.Param(1122334))

			assert.Equal(t, "SELECT * FROM members WHERE (id = $1 AND deleted_at IS NULL)", q.String())
			assert.Equal(t, []interface{}{1122334}, q.Parameters())
		})

		t.Run("in multiple calls", func(t *testing.T) {
			q := From("members")
			q = q.Where("id = $1", q.Param(1122334)).
				Where("deleted_at IS NULL")

			assert.Equal(t, "SELECT * FROM members WHERE (id = $1) AND (deleted_at IS NULL)", q.String())
			assert.Equal(t, []interface{}{1122334}, q.Parameters())
		})

		t.Run("use \"And\" operator", func(t *testing.T) {
			q := From("members")
			q = q.Where("id = $1", q.Param(1122334), q.And("age >= $2", q.Param(18)))

			assert.Equal(t, "SELECT * FROM members WHERE (id = $1 AND (age >= $2))", q.String())
			assert.Equal(t, []interface{}{1122334, 18}, q.Parameters())
		})

		t.Run("use \"Or\" operator", func(t *testing.T) {
			q := From("members")
			q = q.Where("id = $1", q.Param(1122334), q.Or("id = $2", q.Param(1122335)))

			assert.Equal(t, "SELECT * FROM members WHERE (id = $1 OR (id = $2))", q.String())
			assert.Equal(t, []interface{}{1122334, 1122335}, q.Parameters())
		})

		t.Run("nested conditions", func(t *testing.T) {
			q := From("members")
			q = q.Where("A", q.And("B", q.Or("C"), q.Or("D", q.Or("E")))).
				Where("F", q.Or("G")).
				Where("H").
				Where("I", q.Or("J"), q.And("K"))

			assert.Equal(t, "SELECT * FROM members WHERE (A AND (B OR (C) OR (D OR (E)))) AND (F OR (G)) AND (H) AND (I OR (J) AND (K))", q.String())
		})
	})
}
