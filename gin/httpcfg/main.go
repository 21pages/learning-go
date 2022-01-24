package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	s := http.Server{
		Addr:           ":8080",
		Handler:        r,
		MaxHeaderBytes: 100,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}
	s.ListenAndServe()
}
