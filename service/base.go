package service

import (
	"network/utils/inject"
	"reflect"
)

//srv 基础包

type Srv struct {
	Start IStartSrv `auto:"start"`
	App   IAppSrv   `auto:"app"`
}

func (srv *Srv) SrvInject() *Srv {
	inject.Register("start", &Start{})
	inject.Register("app", &App{})

	inject.AutoRegister(srv)

	inject.Inject()

	value := reflect.ValueOf(srv)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	for i := 0; i < value.NumField(); i++ {
		name := value.Type().Field(i).Tag.Get("auto")
		// 获取到tag auto 的值
		switch name {
		case "start":
			srv.Start = &Start{}
		case "app":
			srv.App = &App{}
		}
	}
	return srv
}
