package api

import (
	"fmt"

	"github.com/Alifarid0011/shop/src/api/routers"
	"github.com/Alifarid0011/shop/src/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		testRouter := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(testRouter)
	}
	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))
	if err != nil {
		return
	}
}
