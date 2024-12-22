package repo

import (
	"Quotium/internal/server/db"
	"Quotium/internal/server/db/entity"
	"github.com/sirupsen/logrus"
)

func GetAllTeachers() []entity.Teacher {
	var teachers []entity.Teacher
	if err := db.DB().Find(&teachers).Error; err != nil {
		logrus.Errorf("Failed to retrieve teachers from db: %v", err)
		return nil
	}
	return teachers
}

func AddTeachers(teachers []entity.Teacher) bool {
	if err := db.DB().Create(&teachers).Error; err != nil {
		logrus.Errorf("Failed to add teachers to db: %v", err)
		return false
	}
	return true
}
