package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/d-cryptic/crm-golang-backend/controllers"
	"github.com/d-cryptic/crm-golang-backend/middleware"
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
	}

	return router
}