package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/AbdulrahmanDaud10/jwtethereumwalletauthentication"
	"github.com/AbdulrahmanDaud10/jwtethereumwalletauthentication/pkg/api"
)

func main() {
	app := jwtethereumwalletauthentication.Init()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET, POST, OPTIONS, PUT, DELETE"},
		AllowHeaders: []string{"*"},
	}))

	r.Use(func(c *gin.Context) {
		c.Set("app", app)
	})

	core := r.Group("/core")
	{
		core.GET("/gasPrice", api.GetGasPrice)
		core.GET("/balance/:address", api.GetBalance)
	}

	auth := r.Group("/auth")
	{
		auth.GET("/nonce/:address", app.Nonce)
		auth.POST("/signin", app.Signin)
	}

	users := r.Group("/users").Use(app.Auth())
	{
		users.GET("/me", api.GetUser)
	}

	r.Run()
}
