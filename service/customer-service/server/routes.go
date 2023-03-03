package server

import (
	"customer-service/customer"

	"github.com/gin-gonic/gin"
)

func customerRoutes(r *gin.Engine, controller customer.CustomerController){
	customer := r.Group("/customer")
	customer.POST("/register", controller.CustomerRegister)
	customer.GET("/:id", controller.GetUser)
	customer.POST("/limit", controller.SetCustomerLimit)
	customer.PATCH("/limit/:customer_id/:tenor", controller.UpdateLimit)
}