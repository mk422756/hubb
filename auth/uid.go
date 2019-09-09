package auth

import (
	"context"
)

type contextKey string

var (
	uid contextKey = "uid"
)

func SetUID(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, uid, value)
}

func GetUID(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(uid).(string)
	return val, ok
}
