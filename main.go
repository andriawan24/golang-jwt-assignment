package main

import (
	"assignment-3/config"
	"assignment-3/controllers"
	"assignment-3/repositories"
	"assignment-3/routes"
	"assignment-3/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("File .env not found")
	}

	dbHost := os.Getenv("DB_HOST")
	db, err := config.GetDatabase(dbHost)
	if err != nil {
		log.Fatal("Failed to connect to database", err.Error())
	}

	orderRepository := repositories.NewOrderRepository(db)
	itemRepository := repositories.NewItemRepository(db)
	orderService := services.NewOrderService(orderRepository, itemRepository)
	orderController := controllers.NewOrdersController(orderService)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService()
	userController := controllers.NewUserController(userService, authService)

	router := gin.Default()
	// Order Routes
	routes.OrderRoutes(router, authService, orderController)
	// Authentication Routes
	routes.AuthRoutes(router, authService, userController)

	router.Run(":" + os.Getenv("SERVICE_PORT"))
}
