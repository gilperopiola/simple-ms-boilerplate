package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterActions interface {
	Setup()
}

type MyRouter struct {
	*gin.Engine
}

func (router *MyRouter) Setup() {
	gin.SetMode(gin.DebugMode)
	router.Engine = gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Authentication", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Authentication", "Authorization", "Content-Type"},
	}))

	v1 := router.Group("/v1")
	{
		v1.POST("/Ping", Ping)
	}
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}
