package config

import (
	"context"
	"log"

	"firebase.google.com/go/v4"
)

func InitFirebase() error {
	_, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return nil
}
