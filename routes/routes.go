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

	return router
}