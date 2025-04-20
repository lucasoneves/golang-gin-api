package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasoneves/api-go-gin/controllers"
)

func HandleRoutesRequests() {
	r := gin.Default()
	r.GET("/greeting", controllers.Greeting)
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/students/:id", controllers.GetSingleStudent)
	r.POST("students", controllers.CreateNewStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/search/:cpf", controllers.SearchStudentByCPF)
	r.Run(":5000")
}
