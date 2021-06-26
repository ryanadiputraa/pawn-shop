package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/pawn-shop/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		cookie, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"error": "unauthorized, no cookie found",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// check if request is require admin or employees access
		adminRequest := strings.HasPrefix(c.Request.URL.String(), "/api/employees")
		if adminRequest {
			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(config.GetSecretKey()), nil
			})
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": http.StatusUnauthorized,
					"error": "unauthorized",
				})
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			isAdmin := claims.VerifyIssuer("1", true)
			if !isAdmin {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": http.StatusUnauthorized,
					"error": "unauthorized",
				})
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Next()

		} else {
			_, err = jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(config.GetSecretKey()), nil
			})
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": http.StatusUnauthorized,
					"error": "unauthorized",
				})
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Next()
		}
	}
}