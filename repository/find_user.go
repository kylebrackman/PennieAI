package repository

import (
	"PennieAI/config"
	"fmt"
)

func FindUserByFirebaseUID(firebaseUID string) (bool, error) {
	db := config.GetDB()

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE firebase_uid = $1)", firebaseUID).Scan(&exists)

	fmt.Println("Checking if user exists with Firebase UID:", firebaseUID, "Exists:", exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
