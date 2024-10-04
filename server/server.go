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
	cartRepository := repositories.NewCartRepository(dbHandler)

	adminService := services.NewAdminService(adminRepository)
	beverageService := services.NewBeverageService(beverageRepository)
	userService := services.NewUserService(userRepository)
	cartService := services.NewCartService(cartRepository)

	adminController := controllers.NewAdminController(adminService)
	beverageController := controllers.NewBeverageController(beverageService, adminService)
	userController := controllers.NewUserController(userService)
	cartController := controllers.NewCartController(cartService)


	router := gin.Default()

	router.POST("/user/login", userController.UserLogin)
	router.POST("/user/logout", userController.UserLogout)
	router.POST("/user", userController.CreateUser)
	
	router.POST("/admin/login", adminController.AdminLogin)
	router.POST("/admin/logout", adminController.AdminLogout)

	router.GET("/beverages", beverageController.ListBeverages)
	router.POST("/beverages", beverageController.CreateBeverage)
	router.DELETE("/beverages/:id", beverageController.DeleteBeverage)


	router.GET("/cart", cartController.GetCartItems)
	router.PUT("/cart/add", cartController.AddItem)
	router.DELETE("/cart/remove", cartController.RemoveItem)


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
