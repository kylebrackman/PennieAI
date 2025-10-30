package models

import (
	"time"
)

type Patient struct {
	ID              int        `json:"id" db:"id"`
	Name            string     `json:"name" db:"name"`
	PossibleSpecies []string   `json:"possibleSpecies" db:"possible_species"`
	Breed           *string    `json:"breed" db:"breed"`
	Sex             *string    `json:"sex" db:"sex"`
	DateOfBirth     *time.Time `json:"date_of_birth" db:"date_of_birth"`
	Weight          *float64   `json:"weight" db:"weight"`
	Height          *float64   `json:"height" db:"height"`
	Color           *string    `json:"color" db:"color"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at"`
}
