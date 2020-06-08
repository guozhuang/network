package start

import (
	"github.com/gin-gonic/gin"
	"network/service"
)

type IIndexer interface {
	GetName(ctx *gin.Context)
}

var srv *service.Srv

func init() {
	//将需要加载的model进行依赖注入
	service := &service.Srv{}

	srv = service.SrvInject()
}

type Index struct {
}

func (i *Index) GetName(ctx *gin.Context) {
	var message = ctx.Param("msg")

	result := srv.Start.Say(message)

	ctx.JSON(200, result)
}
