package repository

import (
	"database/sql"
	"errors"

	"PennieAI/config"
	"PennieAI/models"
)

var ErrUserNotFound = errors.New("user not found")

func FindUserByFirebaseUID(firebaseUID string) (models.User, error) {
	db := config.GetDB()

	var user models.User

	err := db.QueryRowx(
		"SELECT id, firebase_uid, email FROM users WHERE firebase_uid = $1",
		firebaseUID,
	).StructScan(&user)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, ErrUserNotFound
		}
		return models.User{}, err
	}

	return user, nil
}
