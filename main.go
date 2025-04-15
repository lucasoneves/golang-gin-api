package main

import (
	"github.com/lucasoneves/api-go-gin/database"
	"github.com/lucasoneves/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRoutesRequests()
}
