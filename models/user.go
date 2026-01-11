package models

import (
	"time"
)

type User struct {
	ID           int        `json:"id" db:"id"`
	FirebaseUID  string     `json:"firebaseUID" db:"firebase_uid"`
	Email        string     `json:"email" db:"email"`
	FirstName    *string    `json:"firstName" db:"first_name"`
	LastName     *string    `json:"lastName" db:"last_name"`
	PhotoURL     *string    `json:"photoURL" db:"photo_url"`
	LastSignInAt *time.Time `json:"lastSignInAt" db:"last_sign_in_at"`
	CreatedAt    time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time  `json:"updatedAt" db:"updated_at"`
}
