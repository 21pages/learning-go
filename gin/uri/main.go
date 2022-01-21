package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		var p Person
		if err := c.ShouldBindUri(&p); err != nil {
			c.JSON(400, gin.H{
				"msg": err,
			})
		}
		c.JSON(200, gin.H{
			"id(uuid)": p.ID,
			"name":     p.Name,
		})
	})
	r.Run()
}

/*
测试:
curl -v localhost:8080/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3 (必须得是uuid格式才能行)
curl -v localhost:8080/thinkerou/not-uuid
*/
