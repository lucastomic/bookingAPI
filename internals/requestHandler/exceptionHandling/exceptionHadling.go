package exceptionhandling

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucastomic/naturalYSalvajeRent/internals/exceptions"
)

// HandleException takes a gin.Context pointer (c) and an error (err).
// If the error is an exceptions.ApiError, it will update c with the status code and message of the error.
// Otherwise, it will update c with a 500 Internal Server Error and the message of the error.
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
