package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	engine := gin.Default()

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	engine.Use(func(context *gin.Context) {
		log.Info("incoming request", zap.String("path", context.Request.URL.Path))
		context.Next()
	})

	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/hello", func(context *gin.Context) {
		context.String(200, "world")
	})
	engine.Run()
}
