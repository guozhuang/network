package controller

import (
	"network/controller/start"
	"network/utils/inject"
	"reflect"
)

//公共的控制器意义不大，也就是对应的包的隔离使用【是否有必要？似乎也没有】【就进行一个通用的处理访问处理？】
//因为公共控制器处理的逻辑应该放到中间层来实现【log+鉴权】

func init() {
	//将对应模块进行初始化并且绑定
}

type Controller struct {
	Start start.IIndexer `auto:"start"`
}

func (controller *Controller) ControllerInject() *Controller {
	inject.Register("start", &start.Index{})

	inject.AutoRegister(controller)

	inject.Inject()

	value := reflect.ValueOf(controller)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	for i := 0; i < value.NumField(); i++ {
		name := value.Type().Field(i).Tag.Get("auto")
		// 获取到tag auto 的值
		switch name {
		case "start":
			controller.Start = &start.Index{}
		}
	}
	return controller
}
