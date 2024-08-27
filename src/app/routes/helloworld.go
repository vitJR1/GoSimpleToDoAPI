package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloWorld
// @Summary Hello world
// @Tags Default
// @Router / [get]
// @Accept json
// @Produce json
func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
