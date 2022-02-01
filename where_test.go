package pgxbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhere(t *testing.T) {
	t.Run("with a condition", func(t *testing.T) {
		q := From("posts").
			Where("id = 300")

		assert.Equal(t, "SELECT * FROM posts WHERE (id = 300)", q.String())
	})

	t.Run("with many conditions", func(t *testing.T) {
		t.Run("with many Where() calls", func(t *testing.T) {
			q := From("members").
				Where("id = 1122334").
				Where("deleted_at IS NULL")

			assert.Equal(t, "SELECT * FROM members WHERE (id = 1122334) AND (deleted_at IS NULL)", q.String())
		})

		t.Run("use \"And\" operator)", func(t *testing.T) {
			q := From("members")
			q = q.Where("id = 1122334", q.And("deleted_at IS NULL"))

			assert.Equal(t, "SELECT * FROM members WHERE (id = 1122334 AND (deleted_at IS NULL))", q.String())
		})

		t.Run("use \"Or\" operator)", func(t *testing.T) {
			q := From("members")
			q = q.Where("id = 1122334", q.Or("id = 1122335"))

			assert.Equal(t, "SELECT * FROM members WHERE (id = 1122334 OR (id = 1122335))", q.String())
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
