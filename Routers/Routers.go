package Routers

import (
	"bingchat-gpt4-backend/Common/RestResponse"
	"bingchat-gpt4-backend/Common/Utils"
	"github.com/gin-gonic/gin"
)

func Init(server *gin.Engine) {
	baseGroup := server.Group("/api")
	baseGroup.GET("/ping", func(ctx *gin.Context) {
		RestResponse.Success(ctx, "", "Pong")
	})

	_ = server.Run(Utils.ReadConfigVal("server.port").(string))
}
