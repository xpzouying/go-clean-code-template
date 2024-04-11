package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WriteSimpleSuccessResponse writes a simple success response to the client.
func WriteSimpleSuccessResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}

// WriteSuccessResponse writes a success response to the client.
func WriteSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   data,
	})
}

// WriteErrorResponseByErrorRequest writes an error response to the client for invalid request.
func WriteErrorResponseByErrorRequest(c *gin.Context) {
	WriteErrorResponse(c, "invalid request parameters")
}

// WriteErrorResponse writes an error response to the client.
func WriteErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"result": false,
		"error":  message,
	})
}
