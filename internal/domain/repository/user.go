package repository

import (
	"service-user/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryUser interface {
	RegisterUser(user entity.User) (*entity.User, error)
	UpdateUser(user entity.User) (*entity.User, error)
	FindUserByEmailAndPassword(email string) *entity.User
}

func NewRepositoryUser(db *gorm.DB) IRepositoryUser {
	return &repositoryUser{
		db,
	}
}

type repositoryUser struct {
	db *gorm.DB
}

func (r *repositoryUser) RegisterUser(user entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error
	return &user, err
}

func (r *repositoryUser) UpdateUser(user entity.User) (*entity.User, error) {
	err := r.db.Save(&user).Error
	return &user, err

}

func (r *repositoryUser) FindUserByEmailAndPassword(email string) *entity.User {
	var user entity.User
	r.db.First(&user, "correo = ? ", email)
	return &user
}
