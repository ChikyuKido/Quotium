package quote

import (
	"Quotium/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListQuotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, helper.QuoteList(c))
	}
}
