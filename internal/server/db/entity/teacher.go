package entity

type Teacher struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	Title     string `gorm:"size:255;not null"`
	ShortName string `gorm:"size:5;not null"`
}
