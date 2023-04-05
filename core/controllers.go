package core

import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	scheme := "http"

	if c.Request.TLS != nil {
		scheme = "https"
	}

	c.JSON(200, CreateIndex(scheme+"://"+c.Request.Host))
}
