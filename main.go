package main

import (
	"Quotium/internal/manager"
	"Quotium/internal/server"
	"Quotium/internal/server/db"
	"Quotium/util"
	"github.com/ChikyuKido/wat/wat"
	static "github.com/ChikyuKido/wat/wat/server/static"
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
	_ = os.MkdirAll("data", 0750)
	firstStart := !util.FileExists("data/database.db")
	db.InitDatabase()
	if firstStart {
		manager.UpdateTeachersInDB()
	}
	r := gin.Default()
	static.LoadTemplates("./website/templates")
	wat.Roles["unverifiedUser"] = append(wat.Roles["unverifiedUser"], "index")
	wat.Roles["admin"] = append(wat.Roles["admin"], "index", "createQuote", "listQuotes", "listTeachers")
	wat.Roles["user"] = append(wat.Roles["user"], "index", "createQuote", "listQuotes", "listTeachers")

	wat.InitWat(r, db.DB(), firstStart)
	wat.InitWatWebsite(r, "./external/wat/website")
	server.StartServer(r, 8080)
}
