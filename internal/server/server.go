package server

import (
	"Quotium/internal/server/route/quote"
	"Quotium/internal/server/route/sites"
	"Quotium/internal/server/route/teacher"
	wat "github.com/ChikyuKido/wat/wat/server/middleware"
	"github.com/ChikyuKido/wat/wat/server/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

func StartServer(r *gin.Engine, port int) {
	logrus.Info("Starting server on port: " + strconv.Itoa(port))
	initRoutes(r)
	r.Run(":" + strconv.Itoa(port))
}

func initRoutes(r *gin.Engine) {
	quoteRoute := r.Group("/api/v1/quote")
	quoteRoute.POST("/create", wat.RequiredPermission("createQuote", false), quote.CreateQuote())
	quoteRoute.GET("/list", wat.RequiredPermission("listQuotes", false), wat.AuthMiddleware(), quote.ListQuotes())
	teacherRoute := r.Group("/api/v1/teacher")
	teacherRoute.GET("/list", wat.RequiredPermission("listTeachers", false), teacher.ListTeacher())
	teacherRoute.Static("/image", "data/teacher/")
	sitesGroup := r.Group("/")
	sitesGroup.Use(wat.AuthMiddleware())
	static.ServeFolder("/imgs/", "./website/imgs", nil, "imgs", sitesGroup, "notGuest")
	static.ServeFolder("/js/", "./website/js", nil, "js", sitesGroup, "")
	sites.Quotes(sitesGroup)
	sites.Teachers(sitesGroup)
	sites.CreateQuote(sitesGroup)
	sitesGroup.GET("/", wat.RequiredPermission("notGuest", true), static.ServeFile("./website/html/index.html", nil, "quotes"))

}
