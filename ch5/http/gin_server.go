package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const keyRequestId = "requestID"

func main() {
	engine := gin.Default()

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	engine.Use(func(context *gin.Context) {
		s := time.Now()
		context.Next()
		log.Info("incoming request",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("time spent", time.Now().Sub(s)))
	}, func(context *gin.Context) {
		context.Set(keyRequestId, rand.Int())
		context.Next()
	})

	engine.GET("/ping", func(context *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if requestId, exists := context.Get(keyRequestId); exists {
			h[keyRequestId] = requestId
		}
		context.JSON(200, h)
	})

	engine.GET("/hello", func(context *gin.Context) {
		context.String(200, "world")
	})
	engine.Run()
}
