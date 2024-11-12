package routers

import (
	"address-book-backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	applyCORS(router)

	router.POST("/address-books", controllers.CreateAddressBook)
	router.GET("/address-books", controllers.GetAddressBooks)
	router.GET("/address-book/:id", controllers.GetAddressBook)
	router.PUT("/address-book/:id", controllers.UpdateAddressBook)
	router.DELETE("/address-book/:id", controllers.DeleteAddressBook)
	return router
}

func applyCORS(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Authorization", "Accept", "Accept-Version", "Origin"}
	router.Use(cors.New(config))
}
