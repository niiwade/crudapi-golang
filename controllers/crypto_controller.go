package controllers

import (
	"crypto-api/cryptoapi"
	"crypto-api/models"
	"crypto-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CryptoController struct {
	cryptoService services.CryptoService
}

func NewCryptoController(cryptoService services.CryptoService) *CryptoController {
	return &CryptoController{cryptoService}
}

func (ctrl *CryptoController) CreateCryptoRecord(c *gin.Context) {
	var crypto models.CryptoPrice
	if err := c.ShouldBindJSON(&crypto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.cryptoService.CreateData(&crypto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, crypto)
}

func (ctrl *CryptoController) GetCryptoRecord(c *gin.Context) {
	id := c.Param("id") // Get the crypto record ID from the URL parameter
	crypto, err := ctrl.cryptoService.GetData(&id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, crypto)
}

func (ctrl *CryptoController) GetAllCryptoRecords(c *gin.Context) {
	cryptos, err := ctrl.cryptoService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cryptos)
}

func (ctrl *CryptoController) UpdateCryptoRecord(c *gin.Context) {
	var updatedCrypto models.CryptoPrice
	if err := c.ShouldBindJSON(&updatedCrypto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.cryptoService.UpdateData(&updatedCrypto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCrypto)
}

func (ctrl *CryptoController) DeleteCryptoRecord(c *gin.Context) {
	id := c.Param("id") // Get the crypto record ID from the URL parameter
	if err := ctrl.cryptoService.DeleteData(&id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (ctrl *CryptoController) GetCryptoDataController(c *gin.Context) {
	// Call the GetExternalCryptoData function from cryptoapi package
	cryptodata, err := cryptoapi.GetExternalCryptoData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Pass the crypto data to services for further processing

	// Return the crypto data as JSON
	c.JSON(http.StatusOK, cryptodata)
}
