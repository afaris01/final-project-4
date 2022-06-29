package middlewares

import (
	"net/http"
	"strings"
	"final-project-4/helpers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helpers.APIResponse("Unauthorized", "Token Not Found!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Bearer tokentokentoken
		arrayToken := strings.Split(authHeader, " ")

		var tokenString string
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := helpers.NewService().ValidateToken(tokenString)

		if err != nil {
			response := helpers.APIResponse("Unauthorized", "Token Invalid!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helpers.APIResponse("Unauthorized", "Token invalid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["id_user"].(float64))

		if userID == 0 {
			response := helpers.APIResponse("Unauthorized", "ID Not Found!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", userID)
	}
}
