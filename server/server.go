package server

import (
	"database/sql"
	"log"
	"beverages-booking/controllers"
	"beverages-booking/repositories"
	"beverages-booking/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	adminController *controllers.AdminController
	beverageController *controllers.BeverageController
	userController *controllers.UserController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	adminRepository := repositories.NewAdminRepository(dbHandler)
	beverageRepository := repositories.NewBeverageRepository(dbHandler)
	userRepository := repositories.NewUserRepository(dbHandler)

	adminService := services.NewAdminService(adminRepository)
	beverageService := services.NewBeverageService(beverageRepository)
	userService := services.NewUserService(userRepository)

	adminController := controllers.NewAdminController(adminService)
	beverageController := controllers.NewBeverageController(beverageService)
	userController := controllers.NewUserController(userService)


	router := gin.Default()

	router.POST("/user/login", userController.UserLogin)
	
	router.POST("/admin/login", adminController.AdminLogin)
	router.GET("/admin/beverages", beverageController.ListBeverages)
	router.POST("/admin/beverages", beverageController.CreateBeverage)
	router.DELETE("/admin/beverages/:id", beverageController.DeleteBeverage)

	return HttpServer{
		config:            config,
		router:            router,
		adminController: adminController,
		beverageController: beverageController,
		userController: userController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
