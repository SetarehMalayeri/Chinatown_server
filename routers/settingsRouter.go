package routers

import (
	"chinatown_server/api"
)

func (router RouterGroup) SettingsRouter(goup string) {
	settingsRouters := router.Group(goup)
	{
		settingsApi := api.ApiGroupApp.SettingApi
		settingsRouters.GET("/", settingsApi.SettingsInfoView)
	}
}
