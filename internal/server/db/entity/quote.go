package entity

import wat "github.com/ChikyuKido/wat/wat/server/db/entity"

type Quote struct {
	ID           uint     `gorm:"primaryKey"`
	TeacherID    uint     // Foreign-Key
	Teacher      Teacher  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Content      string   `gorm:"type:text;not null"`
	CreationDate string   `gorm:"type:datetime;not null"`
	CreatorID    uint     // Foreign-Key
	Creator      wat.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
