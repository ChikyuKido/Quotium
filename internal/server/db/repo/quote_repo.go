package repo

import (
	"Quotium/internal/server/db"
	"Quotium/internal/server/db/entity"
	"github.com/sirupsen/logrus"
	"time"
)

func CreateQuote(content string, teacherID uint, userID uint, creationDate int64) bool {
	var quote = entity.Quote{
		Content:      content,
		TeacherID:    teacherID,
		CreationDate: creationDate,
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
func GetQuoteCount() int64 {
	var count int64
	if err := db.DB().Model(&entity.Quote{}).Count(&count).Error; err != nil {
		logrus.Errorf("Failed to get quote count %v", err)
		return 0
	}
	return count
}
func GetQuoteCountLast7Days() int64 {
	var count int64
	sevenDaysAgo := time.Now().AddDate(0, 0, -7).Unix()
	if err := db.DB().
		Model(&entity.Quote{}).
		Where("creation_date >= ?", sevenDaysAgo).
		Count(&count).Error; err != nil {
		logrus.Errorf("Failed to get quote count: %v", err)
		return 0
	}
	return count
}
func ListQuotes(teacherID uint, searchQuery string) []entity.Quote {
	var quotes []entity.Quote
	query := db.DB().Preload("Creator").Preload("Teacher")
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
