package sites

import (
	"Quotium/internal/server/db/repo"
	wat "github.com/ChikyuKido/wat/wat/server/middleware"
	static "github.com/ChikyuKido/wat/wat/server/static"
	"github.com/gin-gonic/gin"
)

type stats struct {
	Last7Days     int64
	TotalQuotes   int64
	TotalTeachers int64
}

func Index(r *gin.RouterGroup) {
	r.GET("/", wat.RequiredPermission("index", true), static.ServeFile("./website/html/index.html", func(c *gin.Context) any {
		var data = struct {
			Stats stats
		}{}
		data.Stats.TotalTeachers = repo.GetTeacherCount()
		data.Stats.Last7Days = repo.GetQuoteCountLast7Days()
		data.Stats.TotalQuotes = repo.GetQuoteCount()

		return data
	}, "quotes"))
}
