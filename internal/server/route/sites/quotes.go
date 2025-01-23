package sites

import (
	"Quotium/internal/helper"
	"Quotium/internal/server/db/repo"
	"github.com/ChikyuKido/wat/wat/server/middleware"
	static "github.com/ChikyuKido/wat/wat/server/static"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type quoteData struct {
	Creator string
	Content string
	Teacher string
	Date    string
}
type teacherDropdownData struct {
	ID   string
	Name string
}

func Quotes(r *gin.RouterGroup) {
	r.GET("/quotes", wat.RequiredPermission("listQuotes", true), static.ServeFile("./website/html/quotes.html", func(c *gin.Context) any {
		var data = struct {
			Quotes   []quoteData
			Teachers []teacherDropdownData
		}{}
		dbQuotes := helper.QuoteList(c)
		for _, dbQuote := range dbQuotes {
			creatorName := "anon"
			if dbQuote.CreatorID != 0 {
				creatorName = dbQuote.Creator.Username
			}
			quote := quoteData{
				Creator: creatorName,
				Content: dbQuote.Content,
				Teacher: dbQuote.Teacher.Name,
				Date:    time.Unix(dbQuote.CreationDate, 0).Format("02.01.2006"),
			}
			data.Quotes = append(data.Quotes, quote)
		}
		dbTeachers := repo.GetAllTeachers()
		for _, dbTeacher := range dbTeachers {
			teacher := teacherDropdownData{
				ID:   strconv.Itoa(int(dbTeacher.ID)),
				Name: dbTeacher.Name + "(" + dbTeacher.ShortName + ")",
			}
			data.Teachers = append(data.Teachers, teacher)
		}
		return data
	}, "quotes"))
}
