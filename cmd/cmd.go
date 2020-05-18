package cmd

import (
	"github.com/gin-gonic/gin"
	"network/router"
)

func App() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	//Dependency Injection & Route Register
	router.Configure(r)

	return r
}
