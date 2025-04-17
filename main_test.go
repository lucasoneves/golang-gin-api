package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucasoneves/api-go-gin/controllers"
)

func SetupRouterTest() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestStudentsRoute(t *testing.T) {
	r := SetupRouterTest()

	r.GET("/students", controllers.ShowAllStudents)

	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Fatalf("Status error: Value received: %d - Expected: %d", response.Code, http.StatusOK)
	}
}
