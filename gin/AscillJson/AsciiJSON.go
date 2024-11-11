package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/Somejson", func(c *gin.Context) {
		data := map[string]interface{}{
			"name": "zhangsan",
			"age":  30,
		}
		c.AsciiJSON(http.StatusOK, data)
	})
	r.Run(":8080")

}
