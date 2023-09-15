package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"crypto-api/services"
	"crypto-api/models"
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
