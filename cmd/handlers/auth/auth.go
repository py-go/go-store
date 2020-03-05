package auth

import (
	"go-store/cmd/middleware"
	"go-store/cmd/utils"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddAuth(router *gin.Engine) {

	auth := router.Group("/auth")
	{
		auth.GET("/login", middleware.EnsureNotLoggedIn(), showLoginPage)
		auth.POST("/login", middleware.EnsureNotLoggedIn(), performLogin)
		auth.GET("/logout", middleware.EnsureLoggedIn(), logout)
		auth.GET("/signup", middleware.EnsureNotLoggedIn(), showSignupPage)
		auth.POST("/signup", middleware.EnsureNotLoggedIn(), signup)
	}
}

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func showSignupPage(c *gin.Context) {

	utils.Render(c, "signup.html", gin.H{
		"title": "Register",
	})

}

func signup(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := registerNewUser(username, password); err == nil {

		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		utils.Render(c, "login-successful.html", gin.H{
			"title": "Successful registration & Login",
		})

	} else {
		utils.RenderError(c, http.StatusBadRequest, "signup.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error(),
		})

	}
}

func logout(c *gin.Context) {

	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func showLoginPage(c *gin.Context) {
	utils.Render(c, "login.html", gin.H{
		"title": "Login",
	})
}

func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if isUserValid(username, password) {

		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		utils.Render(c, "login-successful.html", gin.H{
			"title": "Successful Login",
		})

	} else {
		utils.RenderError(c, http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided",
		})

	}
}
