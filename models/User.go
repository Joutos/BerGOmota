package main

import (
	database "bergomota/db"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TODO: Descobrir como usar isso fora da main. Descobrir local certo para usar a main. Meu deus eu estou sofrendo
func main() {
	db := database.Connect()

	// Chamando as migrações
	db.AutoMigrate(User{})
}
