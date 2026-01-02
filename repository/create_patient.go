package repository

import (
	"PennieAI/config"
	"fmt"
)

func CreatePatient(name string) (int, error) {

	db := config.GetDB()

	_, err := db.Exec("INSERT INTO patients (name) VALUES ($1)", name)
	if err != nil {
		fmt.Println("Error inserting user into database:", err)
		return 0, err
	}
	var patientID int

	// Placeholder logic, real implementation will change to use id
	err = db.QueryRow("SELECT id FROM patients WHERE name=$1 ORDER BY id DESC LIMIT 1", name).Scan(&patientID)
	if err != nil {
		fmt.Println("Error retrieving patient ID:", err)
		return 0, err
	}

	return patientID, nil
}
