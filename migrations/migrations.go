package main

import (
	"bergomota/database"
	"bergomota/models"
)

func main() {
	db := database.Connect()

	// Chamando as migrações
	db.AutoMigrate(
		models.User{},
		models.Ponto{},
		models.Empresa{},
		models.Colaborador{},
	)
}
