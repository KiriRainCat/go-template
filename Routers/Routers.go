package Routers

import (
	"github.com/gin-gonic/gin"
	"go-template/Common/RestResponse"
	"go-template/Common/Utils"
)

func Init(server *gin.Engine) {
	baseGroup := server.Group("/api")
	baseGroup.GET("/ping", func(ctx *gin.Context) {
		RestResponse.Success(ctx, "", "Pong")
	})

	_ = server.Run(Utils.ReadConfigVal("server.port").(string))
}
