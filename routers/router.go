package routers

import (
	"address-book-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/address-books", controllers.CreateAddressBook)
	router.GET("/address-books", controllers.GetAddressBooks)
	router.GET("/address-book/:id", controllers.GetAddressBook)
	router.PUT("/address-book/:id", controllers.UpdateAddressBook)
	router.DELETE("/address-book/:id", controllers.DeleteAddressBook)
	return router
}
