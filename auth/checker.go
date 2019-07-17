package auth

import (
	"context"
	"log"
	"os"

	"github.com/hubbdevelopers/db"
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
