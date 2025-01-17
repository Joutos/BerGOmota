package models

import (
	"time"
)

type Colaborador struct {
	IDColaborador   uint `gorm:"primaryKey"`
	HashEmpresa     string
	HashColaborador string
	CPF             string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
