package routers

import (
	"github.com/Alifarid0011/shop/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()
	r.GET("/", h.Test)
	r.POST("/user/:id", h.GetUser)
	r.GET("/user/get-user-by-user-name/:name", h.GetUserByName)
	r.POST("/user/create-user", h.CreateUser)
	r.GET("/query-string", h.GetQueryString)
	r.GET("/header", h.GetHeader)
	r.POST("/body", h.GetBody)
	r.POST("/multi-file", h.MultiFile)
}
