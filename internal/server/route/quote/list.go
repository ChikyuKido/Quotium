package quote

import (
	"Quotium/internal/server/db/entity"
	"Quotium/internal/server/db/repo"
	"github.com/gin-gonic/gin"
	"net/http"
	sort2 "sort"
	"strconv"
)

func ListQuotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		sortType := "cd"
		order := "asc"
		limit := 25
		var teacher uint = 0
		if sortParam := c.Query("sort"); sortParam != "" {
			sortType = sortParam
		}
		if orderParam := c.Query("order"); orderParam != "" {
			order = orderParam
		}
		if limitParam := c.Query("limit"); limitParam != "" {
			if parsedLimit, err := strconv.Atoi(limitParam); err == nil {
				limit = parsedLimit
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse limit url parameter"})
				return
			}
		}
		if teacherParam := c.Query("teacher"); teacherParam != "" {
			if parsedTeacher, err := strconv.ParseUint(teacherParam, 10, 32); err == nil {
				teacher = uint(parsedTeacher)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse teacher url parameter"})
				return
			}
		}
		quotes := repo.ListQuotes(limit, teacher)
		if sortType == "cd" {
			sort2.Sort(entity.ByCreationDate(quotes))
		}
		if order == "desc" {
			reverseSlice(quotes)
		}

		c.JSON(http.StatusOK, quotes)
	}
}
func reverseSlice[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
