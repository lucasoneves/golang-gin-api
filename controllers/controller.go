package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasoneves/api-go-gin/database"
	"github.com/lucasoneves/api-go-gin/models"
)

func ShowAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func GetSingleStudent(c *gin.Context) {
	name := c.Param("id")
	c.JSON(200, gin.H{
		"id": name,
	})
}

func CreateNewStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}
