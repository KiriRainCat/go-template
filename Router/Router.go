package Router

import (
	"go-template/Common/Util/ConfigUtil"
	"go-template/Controller/UserController"
	"go-template/Model/RestResponse"

	"github.com/gin-gonic/gin"
)

func Init(server *gin.Engine) {
	baseGroup := server.Group("/api")
	baseGroup.GET("/ping", func(ctx *gin.Context) {
		RestResponse.Success(ctx, "", "Pong")
	})

	user := baseGroup.Group("/user")
	{
		user.GET("/:id", UserController.GetUser)
		user.POST("", UserController.AddUser)
	}

	_ = server.Run(ConfigUtil.ReadConfigVal("server.port").(string))
}
