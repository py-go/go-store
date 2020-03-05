package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, templateName string, data gin.H) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}

}

func RenderError(c *gin.Context, status int, templateName string, data gin.H) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(status, data)
	case "application/xml":
		c.XML(status, data)
	default:
		c.HTML(status, templateName, data)
	}

}
