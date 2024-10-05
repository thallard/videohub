package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
