package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getx", GetDataX)
	r.GET("/getd", GetDataD)
	r.Run()
}

type StructA struct {
	FieldA string //`form:"field_a"` //注释了好像也没区别
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

type StructX struct {
	X struct{} `form:"name_x"` // 有 form
}

type StructY struct {
	Y StructX `form:"name_y"` // 有 form
}

type StructZ struct {
	Z *StructZ `form:"name_z"` // 有 form
}

func GetDataB(c *gin.Context) {
	var b StructB
	b.NestedStruct = StructA{FieldA: "aaaa"}
	b.FieldB = "bbbb"
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC
	b.NestedStructPointer = &StructA{FieldA: "hello"}
	b.FieldC = "world"
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

func GetDataX(c *gin.Context) {
	var x StructX
	c.Bind(&x)
	c.JSON(200, gin.H{
		"x": x.X,
	})
}
