package models

import (
	"time"
)

type User struct {
	ID            uint `gorm:"primaryKey"`
	Login         string
	Email         string `gorm:"unique"`
	Senha         string
	Administrador bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
