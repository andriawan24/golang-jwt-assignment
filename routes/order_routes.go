package routes

import (
	"assignment-3/controllers"
	"assignment-3/middleware"
	"assignment-3/services"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(route *gin.Engine, authService services.AuthService, ordersController controllers.OrdersController) {
	orderRoutes := route.Group("/orders")
	{
		orderRoutes.GET("/", ordersController.GetAllOrders)
		orderRoutes.GET("/:id", ordersController.GetOrderByID)
		orderRoutes.POST("/", middleware.Authentication(authService), ordersController.CreateOrder)
		orderRoutes.PUT("/:id", middleware.Authentication(authService), ordersController.UpdateOrder)
		orderRoutes.DELETE("/:id", middleware.Authentication(authService), ordersController.DeleteOrder)
	}
}
