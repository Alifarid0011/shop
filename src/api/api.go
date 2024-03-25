package api

import (
	"fmt"

	"github.com/Alifarid0011/shop/docs"
	"github.com/Alifarid0011/shop/src/api/routers"
	"github.com/Alifarid0011/shop/src/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title golang web api
// @version 3.0
// @description golang web api
// @contact.name Alifarid0011
// @contact.url https://github.com/Alifarid0011
// @securityDefinitions.ApiKeyAuth.description bearer token

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	RegisterRouter(r)
	RegisterSwagger(r, cfg)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))
	if err != nil {
		return
	}
}

func RegisterRouter(r *gin.Engine) {

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
}
func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "3.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
