package teacher

import (
	"Quotium/internal/server/db/entity"
	"Quotium/internal/server/db/repo"
	"github.com/gin-gonic/gin"
	"net/http"
	sort2 "sort"
	"strconv"
)

func ListTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		sortType := "name"
		order := "asc"
		limit := 25
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
		teachers := repo.GetTeachers()
		if sortType == "name" {
			sort2.Sort(entity.ByName(teachers))
		} else if sortType == "shortname" {
			sort2.Sort(entity.ByShortName(teachers))
		} else if sortType == "quotes" {
			sort2.Sort(entity.ByQuoteCount(teachers))
		}
		if order == "desc" {
			reverseSlice(teachers)
		}
		if limit < len(teachers) {
			teachers = teachers[:limit]
		}

		c.JSON(http.StatusOK, teachers)
	}
}

func reverseSlice[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
