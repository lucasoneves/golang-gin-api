package main

import (
	"github.com/lucasoneves/api-go-gin/database"
	"github.com/lucasoneves/api-go-gin/models"
	"github.com/lucasoneves/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.StudentsList = []models.Student{
		{
			Name: "Lucas Neves",
			CPF:  "40118610813",
			Rg:   "363730333",
		},
		{
			Name: "Larissa Pardini",
			CPF:  "11111111111",
			Rg:   "485689769",
		},
	}
	routes.HandleRoutesRequests()
}
