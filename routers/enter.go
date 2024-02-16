package routers

import (
	"chinatown_server/global"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.Engine
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) // 用自己的，避免运行时，产生一大坨打印
	r := gin.Default()

	routers := RouterGroup{r}

	routers.SettingsRouter() //系统配置api

	return r
}
