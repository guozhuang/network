package cmd

import (
	"github.com/gin-gonic/gin"
	"network/router"
	"os"
)

func App() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	//环境变量获取，根据不同的环境确定加载的配置文件目录
	isDebug := os.Getenv("DEBUG")
	if isDebug == "1" {
		//fmt.Println("testing")
	}

	//Dependency Injection & Route Register
	router.Configure(r)

	return r
}
