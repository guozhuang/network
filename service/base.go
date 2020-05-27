package service

import (
	"network/service/app"
	"network/service/start"
)

//srv 基础包
type IService interface {
	//使用标准中转方式来实现对内部包的服务的转发？
	GetSrvModel() *Service
}

type Service struct {
	Start start.Service
	App   app.Service
}

//实现一个工厂方法来对应相应的srv模块
func (srv *Service) GetSrvModel() *Service {
	return srv
}
