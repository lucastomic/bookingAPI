package authenticationcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	exceptionhandling "github.com/lucastomic/naturalYSalvajeRent/internals/requestHandler/exceptionHandling"
	serviceports "github.com/lucastomic/naturalYSalvajeRent/internals/services/ports"
)

const authEndpoint = "auth"
const registerEndpoint = authEndpoint + "/register"
const loginEndpoint = authEndpoint + "/login"

func AddEndpoints(r *gin.Engine) {
	r.POST(registerEndpoint, register)
	r.POST(loginEndpoint, login)
}

var authenticationService = serviceports.NewAuthenticationService()

var userInRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func register(c *gin.Context) {
	if err := c.Bind(&userInRequestBody); err != nil {
		exceptionhandling.HandleException(c, err)
		return
	}
	err := authenticationService.Register(userInRequestBody.Email, userInRequestBody.Password)
	if err != nil {
		exceptionhandling.HandleException(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successful registration",
	})
}

func login(c *gin.Context) {
	if err := c.Bind(&userInRequestBody); err != nil {
		exceptionhandling.HandleException(c, err)
		return
	}
	jwt, err := authenticationService.Login(userInRequestBody.Email, userInRequestBody.Password)
	if err != nil {
		exceptionhandling.HandleException(c, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"token": jwt,
	})
}
