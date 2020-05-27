package start

import (
	"github.com/gin-gonic/gin"
	"network/service"
)

type Index struct {
	Service service.IService `inject:""`
}

//此处方法还可以进行封装
func (i *Index) GetName(ctx *gin.Context) {
	var message = ctx.Param("msg")

	result := i.Service.GetSrvModel().Start.Say(message)
	result += i.Service.GetSrvModel().Start.Speak(message)
	result += i.Service.GetSrvModel().App.Say(message)
	ctx.JSON(200, result)
}
