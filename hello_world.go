package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

const version string = "2.0.1"

// VersionHandler handles incoming requests to /version
// and just returns a simple version number

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	fmt.Println("test")
	return r
}

func main() {
	r := setupRouter()
	r.GET("/version", versionHandler)
	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}

func versionHandler(c *gin.Context) {
	io.WriteString(c.Writer, version)
}
