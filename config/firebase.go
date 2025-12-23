package config

import (
	"context"
	"log"

	"firebase.google.com/go/v4"
)

var FirebaseApp *firebase.App

func InitFirebase() (*firebase.App, error) {
	var err error
	FirebaseApp, err = firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return FirebaseApp, nil
}

func GetFirebaseApp() *firebase.App {
	return FirebaseApp
}
