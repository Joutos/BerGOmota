package models

import (
	"time"
)

type Ponto struct {
	ID        uint `gorm:"primaryKey"`
	IdUsuario string
	CreatedAt time.Time
	UpdatedAt time.Time
}
