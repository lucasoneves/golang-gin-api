package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucasoneves/api-go-gin/controllers"
	"github.com/lucasoneves/api-go-gin/database"
	"github.com/lucasoneves/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRouterTest() *gin.Engine {
	routes := gin.Default()
	return routes
}

func CriaAlunoMock() {
	student := models.Student{Name: "Nome do Aluno Teste", CPF: "12345678956", Rg: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeletaAlunoMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestStudentsRoute(t *testing.T) {
	r := SetupRouterTest()
	r.GET("/greeting", controllers.Greeting)

	req, _ := http.NewRequest("GET", "/greeting", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Should be equal")
	mockDaResposta := `{"message":"Hello World"}`
	respostaBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, mockDaResposta, string(respostaBody))

}

func TestGetAllStudents(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRouterTest()
	r.GET("/students", controllers.ShowAllStudents)

	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Should have the same status code")

	fmt.Println(response.Body)

}
