package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasoneves/api-go-gin/controllers"
)

func HandleRoutesRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/students/:id", controllers.GetSingleStudent)
	r.Run(":5000")
}
