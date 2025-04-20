package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasoneves/api-go-gin/database"
	"github.com/lucasoneves/api-go-gin/models"
)

func Greeting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func ShowAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func GetSingleStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
			"status":  http.StatusNotFound,
		})

		return
	}

	c.JSON(200, student)
}

func CreateNewStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudentData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var aluno models.Student
	id := c.Params.ByName("id")

	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted successfully",
	})
	return
}

func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.First(&student, id)
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudentData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
	return

}

func SearchStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
			"status":  http.StatusNotFound,
		})

		return
	}

	c.JSON(http.StatusOK, student)
	return
}
