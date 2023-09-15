package routes

import (
	"crypto-api/controllers"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, cryptoController *controllers.CryptoController) {
	// Create a new crypto record
	router.POST("/crypto", cryptoController.CreateCryptoRecord)

	// Read a crypto record by ID
	router.GET("/crypto/:id", cryptoController.GetCryptoRecord)

	// Read all crypto records
	router.GET("/crypto", cryptoController.GetAllCryptoRecords)

	// Update a crypto record by ID
	router.PUT("/crypto/:id", cryptoController.UpdateCryptoRecord)

	// Delete a crypto record by ID
	router.DELETE("/crypto/:id", cryptoController.DeleteCryptoRecord)

	router.GET("/health", func(c *gin.Context) {
		log.Println("Health check endpoint accessed")
		c.JSON(http.StatusOK, gin.H{
			"status": "Everything is OK",
		})
	})

}
