package main

import (
	"io/fs"
	"jssip_demo_zycoo/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	root, _ := fs.Sub(web.WWWFiles, "www")
	r.StaticFS("/www", http.FS(root))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
