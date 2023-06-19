package Router

import (
	"go-template/Common/RestResponse"
	"go-template/Common/Util/ConfigUtil"

	"github.com/gin-gonic/gin"
)

func Init(server *gin.Engine) {
	baseGroup := server.Group("/api")
	baseGroup.GET("/ping", func(ctx *gin.Context) {
		RestResponse.Success(ctx, "", "Pong")
	})

	_ = server.Run(ConfigUtil.ReadConfigVal("server.port").(string))
}
