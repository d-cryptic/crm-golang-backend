package routes

import (
	"github.com/d-cryptic/crm-golang-backend/controllers"
	"github.com/d-cryptic/crm-golang-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.POST("/register", controllers.CreateUser)
	router.POST("/login", controllers.LoginUser)

	protected := router.Group("/admin")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", controllers.GetUsers)
		protected.GET("/users/:id", controllers.GetUserByID)
		protected.PUT("/users/:id", controllers.UpdateUser)
		protected.DELETE("/users/:id", controllers.DeleteUser)

		protected.POST("/tickets", controllers.CreateTicket)
		protected.PUT("/tickets/:id/resolve", controllers.ResolveTicket)
		// TODO - create get interactions by customer id endpoint

		protected.POST("/interactions", controllers.ScheduleInteraction)
		protected.GET("/interactions/customer/:customer_id", controllers.GetInteractionsByCustomerID)

		// Report endpoints
    	protected.GET("/reports/customer-interactions", controllers.GenerateCustomerInteractionsReport)

		// Email endpoints
    	protected.POST("/send-email", controllers.SendEmailHandler)
    	protected.GET("/track/open/:trackingID", controllers.TrackOpenHandler)

		// activity notifications
		protected.POST("/notifications", controllers.CreateNotification)
		protected.GET("/notifications", controllers.GetNotifications)
	}

	return router
}