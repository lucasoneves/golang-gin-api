package main

import "github.com/gin-gonic/gin"

func ShowAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Lucas Neves",
	})
}

func main() {
	r := gin.Default()
	r.GET("/students", ShowAllStudents)
	r.Run(":5000")
}
