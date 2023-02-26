package main

import (
	"fmt"

	"github.com/BrunoTumelero/go-api/database"
	"github.com/BrunoTumelero/go-api/models"
	"github.com/BrunoTumelero/go-api/routes"
)

func main() {
	models.Persons = []models.Pesonalites{
		{Id: 1, Name: "Nome", History: "bla bla bla"},
		{Id: 2, Name: "Nome 2", History: "ble ble ble"},
	}
	database.ConnectDb()
	fmt.Println("Iniciando o servidor...")
	routes.HandleResquest()
}
