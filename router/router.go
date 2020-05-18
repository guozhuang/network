package router

import (
	"log"
	controller "network/controller/start"
	repository "network/repository/start"
	service "network/service/start"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
)

func Configure(app *gin.Engine) {
	//todo:后续对这里标准化
	var index controller.Index

	//Injection
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: &index},
		&inject.Object{Value: &repository.StartRepo{}},
		&inject.Object{Value: &service.StartService{}},
	)
	if err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("inject fatal: ", err)
	}

	v1 := app.Group("/")
	{
		v1.GET("/get/:msg", index.GetName)
	}
}
