package router

import (
	//"network/controller"
	"github.com/gin-gonic/gin"
	"network/controller"
)

//如何将这里设计的更加标准化
func Configure(app *gin.Engine) {
	//进行对应路由配置，然后初始化相应的服务，用来针对不同的请求来处理

	//通用的路由配置
	v1 := app.Group("/")
	{
		v1.GET("/get/user/:userId", controller.Ctrs.User.GetUserInfo)
	}
}
