package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var serverDir string
	var port int
	flag.StringVar(&serverDir, "dir", "./", "server root dir")
	flag.IntVar(&port, "port", 8888, "server port")
	flag.Parse()

	// 发布服务
	router := gin.Default()
	router.Static("/assets", serverDir)
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		fmt.Println(c.Request.URL.Scheme)
		fmt.Println(c.Request.URL.Host)
		fmt.Println(c.Request.URL.Path)
		fmt.Println(c.Request.URL.RawPath)
		fmt.Println(c.Request.URL.RawQuery)
		c.String(http.StatusOK, message)
	})

	// 监听并服务于 0.0.0.0:8080
	router.Run(fmt.Sprintf(":%d", port))
}
