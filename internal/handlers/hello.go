package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SayHello godoc
// @Summary      Say hello to someone
// @Description  Returns a greeting message for the given name
// @Tags         hello
// @Produce      json
// @Param        name  path      string  true  "Name to greet"
// @Success      200   {object}  map[string]string
// @Router       /say/{name} [get]
func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Dear X" + name,
	})
}
