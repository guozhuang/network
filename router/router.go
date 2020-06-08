package router

import (
	"network/controller"
	//"network/repository"
	//"network/service"

	"github.com/gin-gonic/gin"
)

//如何将这里设计的更加标准化
func Configure(app *gin.Engine) {
	//todo:后续对这里标准化
	controller := &controller.Controller{}

	//todo:Injection统一处理【控制器的注入】
	index := controller.ControllerInject()

	//通用的路由配置
	v1 := app.Group("/")
	{
		v1.GET("/get/:msg", index.Start.GetName)
	}
}
