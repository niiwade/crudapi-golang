package cryptoapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CryptoPrice struct {
	CoinsCount       int     `json:"coins_count"`
	ActiveMarkets    int     `json:"active_markets"`
	TotalMcap        float64 `json:"total_mcap"`
	TotalVolume      float64 `json:"total_volume"`
	BtcD             string  `json:"btc_d"`
	EthD             string  `json:"eth_d"`
	McapChange       string  `json:"mcap_change"`
	VolumeChange     string  `json:"volume_change"`
	AvgChangePercent string  `json:"avg_change_percent"`
	VolumeAth        int64   `json:"volume_ath"`
	McapAth          float64 `json:"mcap_ath"`
}

func GetAllCryptoData() ([]CryptoPrice, error) {
	res, err := http.Get(fmt.Sprintf("%s/global", baseURL))

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var cryptodata []CryptoPrice
	err = json.NewDecoder(res.Body).Decode(&cryptodata)
	if err != nil {
		return nil, err
	}

	return cryptodata, nil
}
