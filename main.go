package main

import (
	"Quotium/internal/manager"
	"Quotium/internal/server"
	"Quotium/internal/server/db"
	"Quotium/util"
	"github.com/ChikyuKido/wat/wat"
	"github.com/ChikyuKido/wat/wat/server/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	godotenv.Load(".env")
	if os.Getenv("DEBUG") != "" && os.Getenv("DEBUG") == "true" {
		logrus.SetLevel(logrus.DebugLevel)
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	firstStart := !util.FileExists("database.db")
	db.InitDatabase()
	if firstStart {
		manager.UpdateTeachersInDB()
	}
	r := gin.Default()
	static.LoadTemplates("./website/templates")
	wat.Roles["user"] = append(wat.Roles["user"], "createQuote")
	wat.Roles["user"] = append(wat.Roles["user"], "listQuotes")
	wat.Roles["user"] = append(wat.Roles["user"], "listTeachers")
	wat.InitWat(r, db.DB(), firstStart)
	wat.InitWatWebsite(r, "./external/wat/website")
	server.StartServer(r, 8080)
}
