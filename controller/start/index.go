package start

import (
	"github.com/gin-gonic/gin"
	"network/service/start"
)

type Index struct {
	Service start.IStartService `inject:""`
}

//此处方法还可以进行封装
func (i *Index) GetName(ctx *gin.Context) {
	var message = ctx.Param("msg")
	ctx.JSON(200, i.Service.Say(message))
}
