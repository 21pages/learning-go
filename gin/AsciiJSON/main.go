package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"aaa": 123,
			"bbb": "hello",
		}
		c.AsciiJSON(200, data)
	})
	r.Run(":8080")
}
