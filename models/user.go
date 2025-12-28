package models

import (
	"time"
)

type User struct {
	ID           int        `json:"id" repository:"id"`
	FirebaseUID  string     `json:"firebaseUID" repository:"firebase_uid"`
	Email        string     `json:"email" repository:"email"`
	FirstName    *string    `json:"firstName" repository:"first_name"`
	LastName     *string    `json:"lastName" repository:"last_name"`
	PhotoURL     *string    `json:"photoURL" repository:"photo_url"`
	LastSignInAt *time.Time `json:"lastSignInAt" repository:"last_sign_in_at"`
	CreatedAt    time.Time  `json:"createdAt" repository:"created_at"`
	UpdatedAt    time.Time  `json:"updatedAt" repository:"updated_at"`
}
