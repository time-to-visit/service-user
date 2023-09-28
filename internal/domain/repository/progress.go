package repository

import (
	"fmt"
	"service-user/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryProgress interface {
	RegisterProgress(progress entity.Progress) (*entity.Progress, error)
	UpdateProgress(progress entity.Progress) (*entity.Progress, error)
	FindProgressByUser(idUser uint64) ([]entity.Progress, error)
	FindProgressByUserAndCategory(idUser uint64, idPrimary uint64, idSecondary uint64, typeID string) (*entity.Progress, error)
	DeleteProgress(id int) error
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
	fmt.Println(idUser)

	err := r.db.Find(&progress, "user_id = ?", idUser).Error
	fmt.Println(progress)
	fmt.Println(err)
	return progress, err
}

func (r *repositoryProgress) FindProgressByUserAndCategory(idUser uint64, idPrimary uint64, idSecondary uint64, typeID string) (*entity.Progress, error) {
	var progress entity.Progress
	err := r.db.First(&progress, "user_id = ? and primary_id = ? and secondary_id  = ? and type = ? ", idUser, idPrimary, idSecondary, typeID).Error
	return &progress, err
}

func (r *repositoryProgress) DeleteProgress(id int) error {
	return r.db.Delete(entity.Progress{}, id).Error
}
