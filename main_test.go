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
	"github.com/stretchr/testify/assert"
)

func SetupRouterTest() *gin.Engine {
	routes := gin.Default()
	return routes
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
	r := SetupRouterTest()
	r.GET("/students", controllers.ShowAllStudents)

	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Should have the same status code")

	fmt.Println(response.Body)

}
