package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", search_data)
	r.Run()
}

/*
ShouldBind checks the Content-Type to select a binding engine automatically,
Depending the "Content-Type" header different bindings are used:
    "application/json" --> JSON binding
    "application/xml"  --> XML binding
otherwise --> returns an error
It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
It decodes the json payload into the struct specified as a pointer.
Like c.Bind() but this method does not set the response status code to 400 and abort if the json is not valid.
*/

func search_data(c *gin.Context) {
	var p Person
	if err := c.ShouldBind(&p); err == nil {
		log.Println(p.Name)
		log.Println(p.Birthday)
		log.Println(p.Address)
	} else {
		log.Println("no search")
	}
}

/*
测试:
curl -X GET "localhost:8080/test?name=appleboy&address=xyz&birthday=1992-03-15"
*/

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}
