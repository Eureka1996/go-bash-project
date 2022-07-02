package web_server

import "github.com/gin-gonic/gin"

func GetEngine() *gin.Engine {
	engine := gin.Default()
	return engine
}

func GinGet(r *gin.Engine) {
	r.GET("/gintest", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "wufuqiang",
			"age":  20,
		})
	})
}
