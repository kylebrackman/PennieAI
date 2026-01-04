package config

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	var err error
	FirebaseApp, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return FirebaseApp, nil
}

func GetFirebaseApp() *firebase.App {
	return FirebaseApp
}
