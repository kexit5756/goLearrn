package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"foot": "bar",
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":8080")

}
