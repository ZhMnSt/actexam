package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WEB(){
	r :=gin.Default()
	r.GET("/hello",func(ctx *gin.Context) {
		ctx.String(http.StatusOK,"Hello")
	})
	r.Run(":8888")
}

func main() {
	WEB()
}
