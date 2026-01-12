package handlers

import (
	"PennieAI/middleware"
	"fmt"

	"github.com/gin-gonic/gin"

	"PennieAI/repository"
)

func CreatePatient(c *gin.Context) {

	doctor, ok := middleware.GetAuthenticatedUser(c)
	if !ok {
		fmt.Println("ERROR: GetAuthenticatedUser failed - check route middleware configuration")
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Printf("Received: %+v\n", req)

	patientName := req.Name

	patientID, err := repository.CreatePatient(patientName, doctor.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create patient"})
		return
	}

	c.JSON(201, gin.H{
		"message":    "Patient created successfully",
		"patient_id": patientID,
	})

}
