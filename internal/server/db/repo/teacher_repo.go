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
func GetTeacherCount() int64 {
	var count int64
	if err := db.DB().Model(&entity.Teacher{}).Count(&count).Error; err != nil {
		logrus.Errorf("Failed to get teacher count %v", err)
		return 0
	}
	return count
}
func GetTeachers(searchQuery string) []entity.Teacher {
	var teachers []entity.Teacher

	err := db.DB().
		Model(&entity.Teacher{}).
		Select("teachers.*, (SELECT COUNT(*) FROM quotes WHERE teacher_id = teachers.id) AS quote_count").
		Where("name LIKE ?", "%"+searchQuery+"%").
		Scan(&teachers).Error
	if err != nil {
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

func GetTeacherById(id uint) *entity.Teacher {
	var teacher entity.Teacher
	if err := db.DB().First(&teacher, id).Error; err != nil {
		logrus.Errorf("Failed to retrieve teacher from db: %v", err)
		return nil
	}
	return &teacher
}
