package models

import (
	"time"
)

type User struct {
	ID           int        `json:"id" db:"id"`
	FirebaseUID  string     `json:"firebase_uid" db:"firebase_uid"`
	Email        string     `json:"email" db:"email"`
	FirstName    *string    `json:"first_name" db:"first_name"`
	LastName     *string    `json:"last_name" db:"last_name"`
	LastSignInAt *time.Time `json:"last_sign_in_at" db:"last_sign_in_at"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}
