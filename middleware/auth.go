package middleware

import (
	b64 "encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kittanutp/salesrecorder/config"
	"github.com/kittanutp/salesrecorder/database"
	"github.com/kittanutp/salesrecorder/service"
)

func AuthApp() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(s, "Basic ")
		str := []string{config.AdminUser, config.AdminPassword}
		if b64.StdEncoding.EncodeToString([]byte(strings.Join(str, ":"))) != token {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token"})
			return
		}
		c.Next()
	}

}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.ValidateToken(tokenString)

		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			username := claims["name"].(string)
			user, err := service.GetUserUsername(database.Connect(), username)
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Set("user", user)
		} else {
			c.AbortWithStatusJSON(401, err)
		}

	}
}
