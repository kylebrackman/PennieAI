package repository

import (
	"PennieAI/config"
	"PennieAI/models"
	"fmt"
)

func CreateUser(user *models.User) error {
	db := config.GetDB()
	_, err := db.Exec("INSERT INTO users (firebase_uid, email, photo_url) VALUES ($1, $2, $3)", user.FirebaseUID, user.Email, user.PhotoURL)
	if err != nil {
		fmt.Println("Error inserting user into database:", err)
		return err
	}
	return nil
}
