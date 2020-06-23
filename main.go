package main

import (
	"network/cmd"
	"network/config"
	"network/controller"
	"network/editor"
	"network/model"
	"network/provider"
)

//var Container *di.Container

func init() {
	//Container = di.NewContainer()//todo:后续改造成基于容器的分层结构来进行项目开发
	//然后再参考现有框架的实现来优化整体代码

	config.Setup()

	editor.Init()
	controller.Init()
	model.Init()
	provider.Init()
	//例如cc配置的初始化和持续刷新的逻辑也放在这里【并且提供相应的接口进行访问】
}

func main() {
	app := cmd.App()
	app.Run()
}
