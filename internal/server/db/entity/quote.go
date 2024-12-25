package entity

import (
	wat "github.com/ChikyuKido/wat/wat/server/db/entity"
)

type Quote struct {
	ID           uint     `gorm:"primaryKey"`
	TeacherID    uint     // Foreign-Key
	Teacher      Teacher  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Content      string   `gorm:"type:text;not null"`
	CreationDate int64    `gorm:"not null"`
	CreatorID    uint     // Foreign-Key
	Creator      wat.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ByCreationDate []Quote

func (q ByCreationDate) Len() int           { return len(q) }
func (q ByCreationDate) Less(i, j int) bool { return q[i].CreationDate < q[j].CreationDate }
func (q ByCreationDate) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
