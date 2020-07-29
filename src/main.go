package main

import (
	"github.com/gin-gonic/gin"
)

func SetRoute(e *gin.Engine){
	e.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{"message": "PONG"})
	})
}

func main(){
	e := gin.Default()
	SetRoute(e)
	e.Run()
}