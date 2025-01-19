package quote

import (
	"Quotium/internal/server/db/repo"
	"github.com/ChikyuKido/wat/wat/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateQuote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData = struct {
			Content      string `json:"content"`
			TeacherID    uint   `json:"teacher_id"`
			CreationDate int64  `json:"creation_date"`
			Anon         bool   `json:"anon"`
		}{}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		if requestData.Content == "" || requestData.TeacherID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		if len(requestData.Content) > 2048 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "content is too long. Max is 2048 characters"})
			return
		}
		date := time.Unix(requestData.CreationDate, 0)
		if date.After(time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "date is in the future"})
			return
		}

		if repo.GetTeacherById(requestData.TeacherID) == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no teacher by this id found"})
			return
		}

		var userID uint = 0
		if !requestData.Anon {
			user := wat.GetUserFromContext(c)
			userID = user.ID
		}

		if !repo.CreateQuote(requestData.Content, requestData.TeacherID, userID, requestData.CreationDate) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create the quote. Try again later"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "successfully created the quote"})
	}
}
