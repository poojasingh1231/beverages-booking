package server

import (
	"beverages_booking/config"
	"beverages_booking/controllers"
	"github.com/gin-gonic/gin"
	"net/http" 
)

func Start() {
	config.InitDB()

	r := gin.Default()

	r.POST("/admin/login", controllers.AdminLogin)
	r.GET("/admin/login-form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/admin/beverages", controllers.ListBeverages)      
	r.POST("/admin/beverages", controllers.CreateBeverage)       
	r.DELETE("/admin/beverages/:id", controllers.DeleteBeverage)

	r.Run(":8080")
}
