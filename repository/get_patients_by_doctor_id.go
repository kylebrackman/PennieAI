package repository

import (
	"errors"

	"PennieAI/config"
	"PennieAI/models"
)

var ErrNoPatientsFound = errors.New("no patients found for the given doctor ID")

func GetPatientsByDoctorID(doctorID int) ([]models.Patient, error) {
	db := config.GetDB()

	var patients []models.Patient
	err := db.Select(&patients, "SELECT id, name, doctor_id FROM patients WHERE doctor_id = $1", doctorID)
	if err != nil {
		return nil, err
	}

	if len(patients) == 0 {
		return nil, ErrNoPatientsFound
	}

	return patients, nil
}
