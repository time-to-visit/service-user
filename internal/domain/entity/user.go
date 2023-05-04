package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Model
	Nombres       string `gorm:"column:nombres;type:varchar(255);not null" json:"nombres" validate:"required"`
	Apellidos     string `gorm:"column:apellidos;type:varchar(255);not null"  json:"apellidos" validate:"required"`
	NombreUsuario string `gorm:"column:nombre_usuario;type:varchar(255);not null"  json:"nombre_usuario" validate:"required"`
	Correo        string `gorm:"column:correo;type:varchar(255);not null; unique"  json:"correo" validate:"required"`
	Clave         string `gorm:"column:clave;type:varchar(255);not null"  json:"clave,omitempty" validate:"required"`
	TipoUsuario   string `gorm:"column:tipo_usuario;type:varchar(255);not null"  json:"tipo_usuario" validate:"required"`
	UrlImagen     string `gorm:"column:url_imagen;type:varchar(255);not null"  json:"-" default:"no tiene"`
}

type UserWihoutValidate struct {
	Model
	Nombres       string `gorm:"column:nombres;type:varchar(255);not null" json:"nombres" validate:"required"`
	Apellidos     string `gorm:"column:apellidos;type:varchar(255);not null"  json:"apellidos" validate:"required"`
	NombreUsuario string `gorm:"column:nombre_usuario;type:varchar(255);not null"  json:"nombre_usuario" validate:"required"`
	Correo        string `gorm:"column:correo;type:varchar(255);not null; unique"  json:"correo" validate:"required"`
	Clave         string `gorm:"column:clave;type:varchar(255);not null"  json:"clave,omitempty" validate:"required"`
	TipoUsuario   string `gorm:"column:tipo_usuario;type:varchar(255);not null"  json:"tipo_usuario" validate:"required"`
	UrlImagen     string `gorm:"column:url_imagen;type:varchar(255);not null"  json:"-" default:"no tiene"`
}

func (m UserWihoutValidate) TableName() string {
	return "users"
}

func (m User) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m User) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
