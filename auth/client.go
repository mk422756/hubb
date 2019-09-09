package auth

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	firebaseAuth "firebase.google.com/go/auth"
)

type Client interface {
	isAlreadyRegisteredUID(ctx context.Context, uid string) bool
	verifyIDToken(ctx context.Context, idToken string) (*firebaseAuth.Token, error)
}

type FirebaseClient struct {
	Client *firebaseAuth.Client
}

func NewClient() (Client, error) {
	if os.Getenv("APP_MODE") != "production" {
		log.Println("using mock authentication client")
		return MockClient{}, nil
	}

	client, err := createFirebaseClient()
	if err != nil {
		log.Printf("error initializing app: %v\n", err)
		return nil, err
	}

	return FirebaseClient{Client: client}, nil
}

func createFirebaseClient() (*firebaseAuth.Client, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (fb FirebaseClient) isAlreadyRegisteredUID(ctx context.Context, uid string) bool {
	_, err := fb.Client.GetUser(ctx, uid)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func (fb FirebaseClient) verifyIDToken(ctx context.Context, idToken string) (*firebaseAuth.Token, error) {
	return fb.Client.VerifyIDToken(ctx, idToken)
}

type MockClient struct {
}

func (MockClient) isAlreadyRegisteredUID(ctx context.Context, uid string) bool {
	return true
}

func (MockClient) verifyIDToken(ctx context.Context, idToken string) (*firebaseAuth.Token, error) {
	return nil, nil
}