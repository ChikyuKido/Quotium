package server

import (
	"Quotium/internal/server/route/quote"
	"Quotium/internal/server/route/sites"
	"Quotium/internal/server/route/teacher"
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
	quoteRoute.POST("/create", quote.CreateQuote())
	quoteRoute.GET("/list", quote.ListQuotes())
	teacherRoute := r.Group("/api/v1/teacher")
	teacherRoute.GET("/list", teacher.ListTeacher())
	teacherRoute.Static("/image", "data/teacher/")
	r.Static("/imgs", "./website/imgs")
	sitesGroup := r.Group("/")
	sites.Quotes(sitesGroup)
	sites.Teachers(sitesGroup)

}
