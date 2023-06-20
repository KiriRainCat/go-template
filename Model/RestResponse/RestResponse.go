package RestResponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    any    `json:"data,omitempty"`
}

func Success(ctx *gin.Context, msg string, data any) {
	if msg == "" {
		msg = "success"
	}
	ctx.JSON(http.StatusOK, &Response{
		Success: true,
		Code:    http.StatusOK,
		Msg:     msg,
		Data:    data,
	})
}

func Failure(ctx *gin.Context, msg string) {
	if msg == "" {
		msg = "failure"
	}
	ctx.JSON(http.StatusOK, &Response{
		Success: false,
		Code:    http.StatusBadRequest,
		Msg:     msg,
		Data:    nil,
	})
}

func ArgumentError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, &Response{
		Success: false,
		Code:    http.StatusBadRequest,
		Msg:     "argument_error",
		Data:    nil,
	})
}
