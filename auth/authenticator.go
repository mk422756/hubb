package auth

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
)

type Authenticator interface {
	IsValidUID(ctx context.Context, uid string) bool
	IsAlreadyRegisteredUID(ctx context.Context, uid string) bool
	UIDMiddleware(h http.Handler) http.Handler
}

type FirebaseAuthenticator struct {
	Client Client
}

func NewAuthenticator(client Client) (Authenticator, error) {
	if os.Getenv("APP_MODE") != "production" {
		log.Println("using Blank Authenticator")
		return BlankAuthenticator{}, nil
	}

	return FirebaseAuthenticator{Client: client}, nil
}

func (fb FirebaseAuthenticator) IsValidUID(ctx context.Context, uid string) bool {

	contextUID, ok := GetUID(ctx)

	if !ok {
		log.Println("Error Auth Checker")
		return false
	}

	if uid == contextUID {
		return true
	}

	log.Println("Auth Check is Invalid.")
	log.Println("user uid: " + uid)
	log.Println("requested uid: " + contextUID)

	return false
}

func (fb FirebaseAuthenticator) IsAlreadyRegisteredUID(ctx context.Context, uid string) bool {
	return fb.Client.isAlreadyRegisteredUID(ctx, uid)
}

func (fb FirebaseAuthenticator) UIDMiddleware(h http.Handler) http.Handler {
	ctx := context.Background()

	fn := func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		idToken := strings.Replace(bearer, "Bearer ", "", 1)

		if idToken != "" {
			token, e := fb.Client.verifyIDToken(r.Context(), idToken)
			if e != nil {
				println(e.Error())
			}

			if token != nil {
				ctx = SetUID(ctx, token.UID)
				r = r.WithContext(ctx)
			}
		}

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

type BlankAuthenticator struct{}

func (fb BlankAuthenticator) IsValidUID(ctx context.Context, uid string) bool {
	return true
}

func (fb BlankAuthenticator) IsAlreadyRegisteredUID(ctx context.Context, uid string) bool {
	return true
}

func (fb BlankAuthenticator) UIDMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
