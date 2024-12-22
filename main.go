package main

import (
	"Quotium/internal/manager"
	"Quotium/internal/server"
	"Quotium/internal/server/db"
	"Quotium/util"
	"github.com/ChikyuKido/wat/wat"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load(".env")
	logrus.SetLevel(logrus.DebugLevel)
	firstStart := !util.FileExists("database.db")
	db.InitDatabase()
	if firstStart {
		manager.UpdateTeachersInDB()
	}
	r := gin.Default()
	wat.InitWat(r, db.DB(), firstStart)
	wat.InitWatWebsite(r)
	server.StartServer(r, 8080)
}
