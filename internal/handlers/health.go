package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LivenessCheck godoc
// @Summary      Liveness check endpoint
// @Description  Returns 200 OK if the service is alive
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /health/liveness [get]
func LivenessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "alive",
	})
}

// ReadinessCheck godoc
// @Summary      Readiness check endpoint
// @Description  Returns 200 OK if the service is ready to handle requests
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /health/readiness [get]
func ReadinessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ready",
	})
}
