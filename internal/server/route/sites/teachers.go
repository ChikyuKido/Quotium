package sites

import (
	"Quotium/internal/helper"
	"github.com/ChikyuKido/wat/wat/server/static"
	"github.com/gin-gonic/gin"
)

type teacherData struct {
	Name       string
	ShortName  string
	QuoteCount int64
	Title      string
	ID         uint
	HasImage   bool
}

func Teachers(r *gin.RouterGroup) {
	r.GET("/teachers", static.ServeFile("./website/html/teachers.html", func(c *gin.Context) any {
		var data = struct {
			Teachers []teacherData
		}{}
		dbTeachers := helper.TeacherList(c)
		for _, dbTeacher := range dbTeachers {
			teacher := teacherData{
				Name:       dbTeacher.Name,
				ShortName:  dbTeacher.ShortName,
				QuoteCount: dbTeacher.QuoteCount,
				Title:      dbTeacher.Title,
				ID:         dbTeacher.ID,
				HasImage:   dbTeacher.HasImage,
			}
			data.Teachers = append(data.Teachers, teacher)
		}
		return data
	}, "quotes"))
}
