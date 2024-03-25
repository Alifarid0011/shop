package routers

import (
	"github.com/Alifarid0011/shop/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthRequest()
	r.GET("/", handler.Health)
	r.POST("/", handler.HealthPost)
	r.GET("/:id", handler.HealthById)
}
