package RestResponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Success bool
	Code    int
	Msg     string
	Data    any
}

func Success(ctx *gin.Context, msg string, data any) {
	if msg == "" {
		msg = "Success"
	}
	ctx.JSON(http.StatusOK, &Response{
		Success: true,
		Code:    http.StatusOK,
		Msg:     msg,
		Data:    data,
	})
}

func Failure(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, &Response{
		Success: false,
		Code:    http.StatusBadRequest,
		Msg:     msg,
		Data:    nil,
	})
}
