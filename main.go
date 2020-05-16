package main

import (
	"github.com/gin-gonic/gin"
)

/*func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
func main() {
	http.HandleFunc("/", sayHello) //注册URI路径与相应的处理函数
	//log.Println("【默认项目】服务启动成功 监听端口 80")
	er := http.ListenAndServe("0.0.0.0:9000", nil)
	if er != nil {
		log.Fatal("ListenAndServe: ", er)
	}
}*/

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
