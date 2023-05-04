package entity

type Progress struct {
	Model
	Score   string              `gorm:"column:type;score:varchar(255);not null" json:"score" validate:"required"`
	UserID  int                 `gorm:"column:user_id;type:int(11);not null" json:"level_id" validate:"required"`
	User    *UserWihoutValidate `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"users,omitempty"`
	ItemID  int                 `gorm:"column:item_id;type:int(11);not null" json:"item_id" validate:"required"`
	GuideID int                 `gorm:"column:guide_id;type:int(11);not null" json:"guide_id" validate:"required"`
}
