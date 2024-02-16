package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	Success = 0
	Error   = 1
)

func Result(code int, data any, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, ctx *gin.Context) {
	Result(Success, data, msg, ctx)
}

func OkWithData(data any, ctx *gin.Context) {
	Result(Success, data, "成功", ctx)
}

func OkWithMessage(msg string, ctx *gin.Context) {
	Result(Success, map[string]any{}, msg, ctx)
}

func Fail(data any, msg string, ctx *gin.Context) {
	Result(Error, data, msg, ctx)
}

func FailWithMessage(msg string, ctx *gin.Context) {
	Result(Error, map[string]any{}, msg, ctx)
}

func FailWithCode(code int, ctx *gin.Context) {
	msg, ok := ErrorMap[ErrorCode(code)]
	if ok {
		Result(int(code), map[string]any{}, msg, ctx)
		return
	}
	Result(Error, map[string]any{}, "未知错误", ctx)
}
