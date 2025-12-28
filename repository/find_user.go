package repository

import (
	"PennieAI/config"
	"fmt"
)

func FindUserByFirebaseUID(firebaseUID string) (bool, error) {
	db := config.GetDB()

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE firebase_uid = $1)", firebaseUID).Scan(&exists)

	if err != nil {
		fmt.Println("Error checking user existence in database:", err)
		return false, err
	}

	return true, nil
}
