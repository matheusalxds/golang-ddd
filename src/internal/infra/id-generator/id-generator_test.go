package idGenerator

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestIdGenerator(t *testing.T) {
	sut := NewIdGenerator()

	t.Run("should generate uuid v7", func(t *testing.T) {
		id, _ := sut.GenerateUUID()

		require.NotEqual(t, id, uuid.Nil)
	})
}
