package cmd

import (
	"github.com/gin-gonic/gin"
	"network/router"
	"os"
)

func App() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	//环境标准化
	isDebug := os.Getenv("GIN_MODE")
	if isDebug != gin.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	//Dependency Injection & Route Register
	router.Configure(r)

	return r
}
