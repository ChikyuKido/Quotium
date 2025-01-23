package db

import (
	"Quotium/internal/server/db/entity"
	"github.com/onrik/gorm-logrus"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{
		Logger: gorm_logrus.New(),
	})
	if err != nil {
		logrus.Fatalf("failed to connect to database: %v", err)
	}
	err = db.AutoMigrate(entity.Teacher{}, entity.Quote{})
	if err != nil {
		logrus.Fatalf("failed to migrate database: %v", err)
	}
	logrus.Info("Database initialized")
}

func DB() *gorm.DB {
	return db
}
