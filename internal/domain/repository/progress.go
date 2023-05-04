package repository

import (
	"service-user/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryProgress interface {
	RegisterProgress(progress entity.Progress) (*entity.Progress, error)
	UpdateProgress(progress entity.Progress) (*entity.Progress, error)
	FindProgressByUser(idUser uint64) ([]entity.Progress, error)
	FindProgressByUserAndCategory(idUser uint64, idCategory uint64) (*entity.Progress, error)
}

func NewRepositoryProgress(db *gorm.DB) IRepositoryProgress {
	return &repositoryProgress{
		db,
	}
}

type repositoryProgress struct {
	db *gorm.DB
}

func (r *repositoryProgress) RegisterProgress(progress entity.Progress) (*entity.Progress, error) {
	err := r.db.Create(&progress).Error
	return &progress, err
}
func (r *repositoryProgress) UpdateProgress(progress entity.Progress) (*entity.Progress, error) {
	err := r.db.Save(&progress).Error
	return &progress, err
}
func (r *repositoryProgress) FindProgressByUser(idUser uint64) ([]entity.Progress, error) {
	var progress []entity.Progress
	err := r.db.Find(&progress, "user_id = ?", idUser).Error
	return progress, err
}

func (r *repositoryProgress) FindProgressByUserAndCategory(idUser uint64, idCategory uint64) (*entity.Progress, error) {
	var progress entity.Progress
	err := r.db.First(&progress, "user_id = ? and item_id = ?", idUser, idCategory).Error
	return &progress, err
}
