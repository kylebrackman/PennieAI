package handlers

import (
	//"PennieAI/models"
	"fmt"

	"github.com/gin-gonic/gin"
	//"net/http"
	"PennieAI/repository"
)

func CreatePatient(c *gin.Context) {
	// Todo: Implement logic to authenticate that the user who sent the request is in firebase and that the request isn't malicious
	var req struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Printf("Received: %+v\n", req)

	patientName := req.Name

	patientID, err := repository.CreatePatient(patientName)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create patient"})
		return
	}

	c.JSON(201, gin.H{
		"message":    "Patient created successfully",
		"patient_id": patientID,
	})
	
}
