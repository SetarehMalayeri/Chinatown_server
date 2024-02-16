package routers

import (
	"chinatown_server/api"
)

// 系统配置api
func (router RouterGroup) SettingsRouter() {
	settingsRouters := router.Group("/settings")
	{
		settingsApi := api.ApiGroupApp.SettingApi
		settingsRouters.GET("/", settingsApi.SettingsInfoView)
	}
}
