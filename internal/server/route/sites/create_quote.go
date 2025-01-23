package sites

import (
	"Quotium/internal/server/db/repo"
	"github.com/ChikyuKido/wat/wat/server/middleware"
	static "github.com/ChikyuKido/wat/wat/server/static"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateQuote(r *gin.RouterGroup) {
	r.GET("/createQuote", wat.RequiredPermission("createQuote", true), static.ServeFile("./website/html/create_quote.html", func(c *gin.Context) any {
		var data = struct {
			Teachers []teacherDropdownData
		}{}
		dbTeachers := repo.GetAllTeachers()
		for _, dbTeacher := range dbTeachers {
			teacher := teacherDropdownData{
				ID:   strconv.Itoa(int(dbTeacher.ID)),
				Name: dbTeacher.Name + "(" + dbTeacher.ShortName + ")",
			}
			data.Teachers = append(data.Teachers, teacher)
		}
		return data
	}, "teacher"))
}
