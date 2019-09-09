package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidUID(t *testing.T) {
	value := "test value"
	client := MockClient{}
	authenticator := FirebaseAuthenticator{Client: client}

	ctx := context.Background()
	ctx = SetUID(ctx, value)

	assert.True(t, authenticator.IsValidUID(ctx, value))
}

func TestIsaValidUID(t *testing.T) {
	value := "test value"
	client := MockClient{}
	authenticator := FirebaseAuthenticator{Client: client}

	ctx := context.Background()
	ctx = SetUID(ctx, value)
	assert.True(t, authenticator.IsAlreadyRegisteredUID(ctx, value))
}
