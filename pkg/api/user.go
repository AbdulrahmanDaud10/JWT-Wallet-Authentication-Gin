package api

import (
	"net/http"

	"github.com/AbdulrahmanDaud10/jwtethereumwalletauthentication"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	app := c.MustGet("app").(*jwtethereumwalletauthentication.App)
	address := c.MustGet("address")

	var user User
	app.Db.Where("address = ?", address).First(&user)
	c.JSON(http.StatusOK, user)
}
