package UserController

import (
	"go-template/Model/Entity"
	"go-template/Model/RestResponse"
	"go-template/Service/UserService"

	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {
	users, err := UserService.GetUserList()
	if err != nil {
		RestResponse.Failure(ctx, err.Error())
		return
	}
	RestResponse.Success(ctx, "", users)
}

func GetUserByID(ctx *gin.Context) {
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

func UpdateUserByID(ctx *gin.Context) {
	userID := ctx.Param("id")

	var user Entity.User
	if err := ctx.BindJSON(&user); err != nil {
		RestResponse.ArgumentError(ctx)
		return
	}

	affected, err := UserService.UpdateUserByID(userID, user)
	if err != nil {
		RestResponse.Failure(ctx, err.Error())
		return
	}
	RestResponse.Success(ctx, "", affected)
}

func DeleteUserByID(ctx *gin.Context) {
	userID := ctx.Param("id")

	affected, err := UserService.DeleteUserByID(userID)
	if err != nil {
		RestResponse.Failure(ctx, err.Error())
		return
	}
	RestResponse.Success(ctx, "", affected)
}
