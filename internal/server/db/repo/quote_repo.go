package repo

import (
	"Quotium/internal/server/db"
	"Quotium/internal/server/db/entity"
	"github.com/sirupsen/logrus"
	"time"
)

func CreateQuote(content string, teacherID uint, userID uint) bool {
	var quote = entity.Quote{
		Content:      content,
		TeacherID:    teacherID,
		CreationDate: time.Now().Unix(),
	}
	if userID != 0 {
		quote.CreatorID = userID
	}
	if err := db.DB().Create(&quote).Error; err != nil {
		logrus.Errorf("Failed to create quote %v", err)
		return false
	}
	return true
}
func ListQuotes(limit int, teacherID uint, searchQuery string) []entity.Quote {
	var quotes []entity.Quote
	query := db.DB().Preload("Creator").Preload("Teacher").Limit(limit)
	if teacherID != 0 {
		query = query.Where(entity.Quote{TeacherID: teacherID})
	}
	if searchQuery != "" {
		query = query.Where("content LIKE ?", "%"+searchQuery+"%")
	}
	if err := query.Find(&quotes).Error; err != nil {
		logrus.Errorf("Failed to list quotes %v", err)
		return nil
	}

	return quotes
}
