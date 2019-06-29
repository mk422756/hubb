package auth

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func UIDMiddleware(h http.Handler) http.Handler {
	opt := option.WithCredentialsFile(os.Getenv("SECRETS_FILE"))
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	fn := func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		idToken := strings.Replace(bearer, "Bearer ", "", 1)
		token, e := client.VerifyIDToken(r.Context(), idToken)
		if e != nil {
			println(e.Error())
		}

		println(idToken)

		if token != nil {
			println(token.UID)
			ctx2 := SetUID(ctx, token.UID)
			r = r.WithContext(ctx2)
		}

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

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
