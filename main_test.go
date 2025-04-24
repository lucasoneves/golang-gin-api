package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucasoneves/api-go-gin/controllers"
	"github.com/lucasoneves/api-go-gin/database"
	"github.com/lucasoneves/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRouterTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
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

func TestBuscaPorCPF(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRouterTest()
	r.GET("/students/search/:cpf", controllers.SearchStudentByCPF)

	req, _ := http.NewRequest("GET", "/students/search/12345678956", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Should have the same status code")

}

func TestSearchStudentById(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRouterTest()
	r.GET("/students/:id", controllers.GetSingleStudent)
	pathDaBusca := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)

	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var studentMock models.Student

	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, "Nome do Aluno Teste", studentMock.Name)
}

func TestDeleteStudent(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupRouterTest()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	pathDeBusca := "/students/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Should have the same status code")
}

func TestUpdateStudent(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRouterTest()
	r.PATCH("/students/:id", controllers.EditStudent)
	pathDeEdicao := "/students/" + strconv.Itoa(ID)

	studentUpdated := models.Student{Name: "Nome do Aluno Teste Atualizado", CPF: "12345678985", Rg: "123456747"}
	studentJson, _ := json.Marshal(studentUpdated)

	req, _ := http.NewRequest("PATCH", pathDeEdicao, bytes.NewBuffer(studentJson))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var studentResponse models.Student
	json.Unmarshal(response.Body.Bytes(), &studentResponse)

	assert.Equal(t, "12345678985", studentResponse.CPF)
	assert.Equal(t, "Nome do Aluno Teste Atualizado", studentResponse.Name)
	assert.Equal(t, "123456747", studentResponse.Rg)

}
