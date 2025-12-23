package handlers

import (
	//"PennieAI/config"
	//"PennieAI/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {

	fmt.Println("Sign Up context here:", c.Request)
	//var req struct {
	//}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sign Up endpoint - to be implemented",
	})

}
