package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasoneves/api-go-gin/models"
)

func ShowAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": models.StudentsList,
	})
}

func GetSingleStudent(c *gin.Context) {
	name := c.Param("id")
	c.JSON(200, gin.H{
		"id": name,
	})
}
