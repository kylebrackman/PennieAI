package handlers

import (
	"PennieAI/middleware"
	"fmt"

	"github.com/gin-gonic/gin"

	"PennieAI/repository"
)

func GetPatients(c *gin.Context) {
	fmt.Println("GetPatients")
	doctor, ok := middleware.GetAuthenticatedUser(c)
	if !ok {
		fmt.Println("ERROR: GetAuthenticatedUser failed - check route middleware configuration")
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}

	patients, err := repository.GetPatientsByDoctorID(doctor.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch patients"})
		return
	}

	c.JSON(200, patients)

}
