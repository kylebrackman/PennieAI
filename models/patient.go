package models

import (
	"time"
)

type Patient struct {
	ID              int        `json:"id" db:"id"`
	Name            string     `json:"name" db:"name"`
	PossibleSpecies *[]string  `json:"possibleSpecies" db:"possible_species"`
	PossibleBreed   *[]string  `json:"possibleBreed" db:"possible_breed"`
	Sex             *string    `json:"sex" db:"sex"`
	DateOfBirth     *time.Time `json:"dateOfBirth" db:"date_of_birth"`
	Weight          *float64   `json:"weight" db:"weight"`
	Height          *float64   `json:"height" db:"height"`
	Color           *string    `json:"color" db:"color"`
	CreatedAt       time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time  `json:"updatedAt" db:"updated_at"`
	DoctorId        int        `json:"doctorId" db:"doctor_id"`
}
