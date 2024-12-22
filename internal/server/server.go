package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

func StartServer(r *gin.Engine, port int) {
	logrus.Info("Starting server on port: " + strconv.Itoa(port))

	r.Run(":" + strconv.Itoa(port))
}
