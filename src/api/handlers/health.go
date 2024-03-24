package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthRequest struct {
}

func NewHealthRequest() *HealthRequest {
	return &HealthRequest{}
}

func (h *HealthRequest) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "working")
	return
}
func (h *HealthRequest) HealthPost(c *gin.Context) {
	c.JSON(http.StatusOK, "working with post method")
	return
}

func (h *HealthRequest) HealthById(c *gin.Context) {
	c.JSON(http.StatusOK, "working by id :"+fmt.Sprintf("%s", c.Param("id")))
	return
}
