package settings_api

import (
	"chinatown_server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, gin.H{"msg": "xxx2"})

	// res.Ok(map[string]string{}, "xxx", ctx)

	// res.OkWithData(map[string]string{
	// 	"id":   "xxx",
	// 	"name": "张三",
	// }, ctx)

	res.FailWithCode(int(res.SettingsError), ctx)
}
