package entity

import (
	"time"

	"gorm.io/gorm"
)

type Progress struct {
	Model
	Score       string              `gorm:"column:score;type:varchar(255);not null" json:"score" validate:"required"`
	UserID      int                 `gorm:"column:user_id;type:int(11);not null" json:"user_id" validate:"required"`
	User        *UserWihoutValidate `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"users,omitempty"`
	Type        string              `gorm:"column:type;type:varchar(255);not null" json:"type" validate:"required"`
	PrimaryID   int                 `gorm:"column:primary_id;type:int(11);not null" json:"primary_id" validate:"required"`
	SecondaryID int                 `gorm:"column:secondary_id;type:int(11);not null" json:"secondary_id" validate:"required"`
}

func (m Progress) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Progress) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
