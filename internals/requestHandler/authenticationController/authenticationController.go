package authenticationcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	exceptionhandling "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/exceptionHandling"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
)

const authEndpoint = "auth"
const registerEndpoint = authEndpoint + "register"

func AddEndpoints(r *gin.Engine) {
	r.POST(registerEndpoint, register)
}

var authenticationService = serviceports.NewAuthenticationService()

var userInRequestBody struct {
	email    string
	password string
}

func register(c *gin.Context) {
	if err := c.Bind(userInRequestBody); err != nil {
		exceptionhandling.HandleException(c, err)
	}
	err := authenticationService.Register(userInRequestBody.email, userInRequestBody.password)
	if err != nil {
		exceptionhandling.HandleException(c, err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successful registration",
	})
}

func login(c *gin.Context) {
	if err := c.Bind(userInRequestBody); err != nil {
		exceptionhandling.HandleException(c, err)
	}
	jwt, err := authenticationService.Login(userInRequestBody.email, userInRequestBody.password)
	if err != nil {
		exceptionhandling.HandleException(c, err)
	}

	c.JSON(http.StatusAccepted, gin.H{
		"token": jwt,
	})
}
