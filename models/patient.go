package models

import (
	"time"
)

type Patient struct {
	ID              int        `json:"id" repository:"id"`
	Name            string     `json:"name" repository:"name"`
	PossibleSpecies *[]string  `json:"possibleSpecies" repository:"possible_species"`
	PossibleBreed   *[]string  `json:"possibleBreed" repository:"possible_breed"`
	Sex             *string    `json:"sex" repository:"sex"`
	DateOfBirth     *time.Time `json:"dateOfBirth" repository:"date_of_birth"`
	Weight          *float64   `json:"weight" repository:"weight"`
	Height          *float64   `json:"height" repository:"height"`
	Color           *string    `json:"color" repository:"color"`
	CreatedAt       time.Time  `json:"createdAt" repository:"created_at"`
	UpdatedAt       time.Time  `json:"updatedAt" repository:"updated_at"`
}
