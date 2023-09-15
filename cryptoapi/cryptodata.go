package cryptoapi

import (
	"crypto-api/models" // Import the models package
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllCryptoData() ([]models.CryptoPrice, error) {
	res, err := http.Get(fmt.Sprintf("%s/global", baseURL))

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var cryptodata []models.CryptoPrice
	err = json.NewDecoder(res.Body).Decode(&cryptodata)
	if err != nil {
		return nil, err
	}

	return cryptodata, nil
}
