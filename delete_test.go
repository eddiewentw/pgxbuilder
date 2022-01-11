package pgxbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	q := Delete("posts")

	assert.Equal(t, "DELETE FROM posts", q.String())
}
