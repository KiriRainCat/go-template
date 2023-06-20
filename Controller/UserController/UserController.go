package UserController

import (
	"go-template/Model/Entity"
	"go-template/Model/RestResponse"
	"go-template/Service/UserService"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	user, err := UserService.GetUserByID(userID)
	if err != nil {
		RestResponse.Failure(ctx, err.Error())
		return
	}
	RestResponse.Success(ctx, "", user)
}

func AddUser(ctx *gin.Context) {
	var user Entity.User
	if err := ctx.BindJSON(&user); err != nil {
		RestResponse.ArgumentError(ctx)
		return
	}

	id, err := UserService.AddUser(user)
	if err != nil {
		RestResponse.Failure(ctx, err.Error())
		return
	}
	RestResponse.Success(ctx, "", id)
}
