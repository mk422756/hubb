package auth

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/hubbdevelopers/db"
	"google.golang.org/api/option"
)

func Check(ctx context.Context, user db.User) bool {
	if os.Getenv("APP_MODE") == "debug" {
		log.Println("Auth Checker is debug mode")
		return true
	}

	uid, ok := GetUID(ctx)

	if !ok {
		log.Println("Error Auth Checker")
		return false
	}

	if user.UID == uid {
		return true
	}

	log.Println("Auth Check is Invalid.")
	log.Println("user uid: " + user.UID)
	log.Println("requested uid: " + uid)

	return false
}

func CheckValidUID(ctx context.Context, uid string) bool {
	if os.Getenv("APP_MODE") == "debug" {
		log.Println("Auth Checker is debug mode")
		return true
	}

	opt := option.WithCredentialsFile(os.Getenv("SECRETS_FILE"))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	_, err = client.GetUser(ctx, uid)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
