package middleware

import (
	"go-store/cmd/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
			utils.RenderError(c, http.StatusUnauthorized, "error.html", gin.H{
				"message": "401 errors",
				"title":   "Error",
			})

		}
	}
}

func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
			utils.RenderError(c, http.StatusUnauthorized, "error.html", gin.H{
				"message": "401 errors",
				"title":   "Error",
			})
		}
	}
}

func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
