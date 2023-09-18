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


func GetExternalCryptoData() ([]models.CryptoPrice, error) {
    // Define the endpoint URL based on the baseURL
    endpoint := fmt.Sprintf("%s/global", baseURL)

    // Send a GET request to the API endpoint
    res, err := http.Get(endpoint)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    // Check the HTTP status code
    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("HTTP request failed with status code %d", res.StatusCode)
    }

    // Decode the JSON response into your CryptoPrice struct
    var cryptodata []models.CryptoPrice
    err = json.NewDecoder(res.Body).Decode(&cryptodata)
    if err != nil {
        return nil, err
    }

    return cryptodata, nil
}
