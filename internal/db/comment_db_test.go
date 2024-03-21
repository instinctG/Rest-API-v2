//go:build integration
// +build integration

package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentDataBase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), "1e8abc78-f8cf-43ff-a826-176ede6f871d")

		assert.NoError(t, err)
	})
}
