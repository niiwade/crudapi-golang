package services

import "crypto-api/models"

type CryptoService interface {
	CreateData(*models.CryptoPrice) error
	GetData(*string) (*models.CryptoPrice, error)
	GetExternalCryptoData(*string) (*models.CryptoPrice, error)
	GetAll() ([]*models.CryptoPrice, error)
	UpdateData(*models.CryptoPrice) error
	DeleteData(*string) error
}
