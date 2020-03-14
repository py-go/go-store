package middlewares

import (
	"errors"
	"fmt"
	"go-store/cmd/dtos"
	"go-store/cmd/models"
	svc "go-store/cmd/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Add CommonMiddleware

func EnforceAuthenticatedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("storeUser")
		if exists && user.(models.User).ID != 0 {
			return
		} else {
			err := errors.New("Insufficient permission.")
			//Ref: https://godoc.org/github.com/gin-gonic/gin#Context.AbortWithStatusJSON
			c.AbortWithStatusJSON(http.StatusForbidden, dtos.ErrorResponse("access_error", err))
			return
		}
	}
}

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		bearer := c.Request.Header.Get("Authorization")
		if bearer != "" {
			jwtParts := strings.Split(bearer, " ")
			if len(jwtParts) == 2 {
				jwtEncoded := jwtParts[1]

				token, err := jwt.Parse(jwtEncoded, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signin method %v", token.Header["alg"])
					}
					secret := []byte(os.Getenv("JWT_SECRET"))
					return secret, nil
				})

				if err != nil {
					log.Printf("Error while jwt parsing  %v", err)
					return
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					userId := uint(claims["user_id"].(float64))
					var user models.User
					user.ID = userId
					user, err := svc.FindOneUser(&user)
					if err != nil {
						log.Printf("Error while jwt parsing  %v", err)
						return
					}
					c.Set("storeUser", user)
					c.Set("storeUserId", user.ID)

				}

			}
		}
	}
}
