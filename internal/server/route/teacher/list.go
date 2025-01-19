package teacher

import (
	"Quotium/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, helper.TeacherList(c))
	}
}
