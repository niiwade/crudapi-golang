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

	// Use the cryptoService from the controller
	cryptoService := ctrl.cryptoService

	// Call the GetExternalCryptoData function from cryptoapi package
	cryptodata, err := cryptoapi.GetExternalCryptoData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Process the fetched data and insert it into MongoDB
	for _, data := range cryptodata {
		// Create a CryptoPrice instance
		cryptoPrice := &models.CryptoPrice{
			// Map fields from the data to your struct fields
			CoinsCount:       data.CoinsCount,
			ActiveMarkets:    data.ActiveMarkets,
			TotalMcap:        data.TotalMcap,
			TotalVolume:      data.TotalVolume,
			BtcD:             data.BtcD,
			EthD:             data.EthD,
			McapChange:       data.McapChange,
			VolumeChange:     data.VolumeChange,
			AvgChangePercent: data.AvgChangePercent,
			VolumeAth:        data.VolumeAth,
			McapAth:          data.McapAth,
		}

		// Insert the CryptoPrice instance into MongoDB
		if err := cryptoService.CreateData(cryptoPrice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Pass the crypto data to services for further processing

	// Return the crypto data as JSON
	c.JSON(http.StatusOK, cryptodata)
}
