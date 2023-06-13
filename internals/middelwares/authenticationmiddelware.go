package middelwares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	serviceinjector "github.com/lucastomic/naturalYSalvajeRent/internals/services/injection"
)

var authService = serviceinjector.NewAuthenticationService()

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		tokenString = removeBearerHeader(tokenString)
		err := authService.Validate(tokenString)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
func removeBearerHeader(token string) string {
	return strings.Split(token, " ")[1]
}
