package settings_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "xxx2",
	})
}
