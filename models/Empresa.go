package models

import (
	"time"
)

type Empresa struct {
	IDEmpresa   uint `gorm:"primaryKey"`
	Login       string
	Email       string `gorm:"unique"`
	Senha       string
	CNPJ        bool
	HashEmpresa string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
