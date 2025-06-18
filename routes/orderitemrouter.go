package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/ramkrishnareddy24/GoRestaurant/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderitems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderitems/:order_item_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderitems-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderitems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderitems/:order_item_id", controllers.UpdateOrderItem())
}