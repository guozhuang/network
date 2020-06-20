package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"network/config"
	"network/router"
)

func App() *gin.Engine {
	gin.SetMode(config.ServerSetting.RunMode)

	r := gin.New()
	r.Use(gin.Logger()) //gin的中间层的使用
	r.Use(gin.Recovery())

	router.Configure(r)

	endPoint := fmt.Sprintf(":%d", config.ServerSetting.HttpPort)

	//绑定对应的端口
	server := &http.Server{
		Addr:         endPoint,
		Handler:      r,
		ReadTimeout:  config.ServerSetting.ReadTimeout,
		WriteTimeout: config.ServerSetting.WriteTimeout,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

	return r
}
