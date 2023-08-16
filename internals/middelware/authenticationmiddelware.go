package middelware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	serviceinjector "github.com/lucastomic/naturalYSalvajeRent/internals/services/injection"
	authenticationstate "github.com/lucastomic/naturalYSalvajeRent/internals/state/authentication"
)

var jwtService = serviceinjector.NewJWTService()
var authenticationService = serviceinjector.NewAuthenticationService()

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := getAuthorizationHeader(context)
		if tokenDoesntExist(token) {
			setUnexistentTokenMessage(context)
			return
		}
		token = removeBearerHeader(token)
		err := jwtService.Validate(token)
		if err != nil {
			setTokenErrorMessage(context, err)
			return
		}
		user := getUserFromTokenString(token)
		authenticationstate.SetAuthenticatedUser(user)
		context.Next()
		// defer authenticationstate.RemoveAuthenticatedUser()
	}
}

func getUserFromTokenString(token string) domain.User {
	email, _ := jwtService.GetEmail(token)
	user, _ := authenticationService.GetUser(email)
	return user
}

func setTokenErrorMessage(context *gin.Context, err error) {
	context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	context.Abort()
}
func setUnexistentTokenMessage(context *gin.Context) {
	context.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
	context.Abort()
}

func getAuthorizationHeader(context *gin.Context) string {
	return context.GetHeader("Authorization")
}

func tokenDoesntExist(token string) bool {
	return token == ""
}

func removeBearerHeader(token string) string {
	return strings.Split(token, " ")[1]
}
