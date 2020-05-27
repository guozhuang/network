package router

import (
	"log"
	controller "network/controller/start"
	"network/repository"
	"network/service"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
)

//如何将这里设计的更加标准化
func Configure(app *gin.Engine) {
	//todo:后续对这里标准化
	var index controller.Index

	//Injection
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: &index},
		&inject.Object{Value: &repository.Repo{}},
		&inject.Object{Value: &service.Service{}},
	)
	if err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("inject fatal: ", err)
	}

	//通用的路由配置
	v1 := app.Group("/")
	{
		v1.GET("/get/:msg", index.GetName)
	}
}
