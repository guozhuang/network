package main

import (
	"network/cmd"
	"network/config"
	"network/controller"
	"network/model"
	"network/provider"
)

func init() {
	config.Setup()

	controller.Init()
	model.Init()
	provider.Init()
	//例如cc配置的初始化和持续刷新的逻辑也放在这里【并且提供相应的接口进行访问】
	//redis的连接池的初始化
}

func main() {
	// 创建一个默认的路由引擎
	/*r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()*/

	app := cmd.App()
	app.Run()
}
