package UserController

import (
	"go-template/Model/Entity"
	"go-template/Model/RestResponse"
	"go-template/Service/UserService"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		RestResponse.Failure(ctx, "No empty query")
		return
	}

	user := UserService.GetUserByID(userID)
	if user.ID == 0 {
		RestResponse.Failure(ctx, "No result found")
		return
	}
	RestResponse.Success(ctx, "", user)
}

func AddUser(ctx *gin.Context) {
	var user Entity.User
	_ = ctx.BindJSON(&user)
	if user.Username == "" || user.Password == "" {
		RestResponse.Failure(ctx, "No empty query")
		return
	}

	if UserService.AddUser(user) != nil {
		RestResponse.Failure(ctx, "DB insert error")
		return
	}
	RestResponse.Success(ctx, "", user.Username)
}
