package exceptionhandling

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"
)

func HandleException(c *gin.Context, err error) {
	if apiErr, ok := err.(exceptions.ApiError); ok {
		status, msg := apiErr.APIError()
		c.JSON(status, gin.H{
			"error": msg,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

}
