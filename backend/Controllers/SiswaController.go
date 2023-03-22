package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email    string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	siswa := models.Siswa{}

	siswa.Email = input.Email
	siswa.Password = input.Password

}
