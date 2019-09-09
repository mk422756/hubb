package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAndGetUID(t *testing.T) {
	value := "test value"
	ctx := context.Background()
	ctx = SetUID(ctx, value)

	ret, ok := GetUID(ctx)

	assert.Equal(t, ret, value)
	assert.True(t, ok)
}
