package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET, POST, OPTIONS, PUT, DELETE"},
		AllowHeaders: []string{"*"},
	}))

	r.Run()
}
