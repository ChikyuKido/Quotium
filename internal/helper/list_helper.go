package helper

import "C"
import (
	"Quotium/internal/server/db/entity"
	"Quotium/internal/server/db/repo"
	"github.com/gin-gonic/gin"
	"net/http"
	sort2 "sort"
	"strconv"
)

func QuoteList(c *gin.Context) []entity.Quote {
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
		if limitParam == "all" {
			limit = 2_000_000
		} else if parsedLimit, err := strconv.Atoi(limitParam); err == nil {
			limit = parsedLimit
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse limit url parameter"})
			return nil
		}
	}
	if teacherParam := c.Query("teacher"); teacherParam != "" {
		if parsedTeacher, err := strconv.ParseUint(teacherParam, 10, 32); err == nil {
			teacher = uint(parsedTeacher)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse teacher url parameter"})
			return nil
		}
	}
	quotes := repo.ListQuotes(limit, teacher, c.Query("search"))
	if sortType == "cd" {
		sort2.Sort(entity.ByCreationDate(quotes))
	}
	if order == "desc" {
		reverseSlice(quotes)
	}
	return quotes
}

func TeacherList(c *gin.Context) []entity.Teacher {
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
		if limitParam == "all" {
			limit = 2_000_000
		} else if parsedLimit, err := strconv.Atoi(limitParam); err == nil {
			limit = parsedLimit
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse limit url parameter"})
			return nil
		}
	}
	teachers := repo.GetTeachers(c.Query("search"))
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

	return teachers
}

func reverseSlice[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
