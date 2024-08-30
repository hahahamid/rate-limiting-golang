package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMessage handles the GET request to /api/message
func GetMessage(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}
