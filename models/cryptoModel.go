package models
import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CryptoPrice struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	CoinsCount       int     `json:"coins_count" bson:"coins_count"`
	ActiveMarkets    int     `json:"active_markets"  bson:"active_markets"`
	TotalMcap        float64 `json:"total_mcap"  bson:"total_mcap"`
	TotalVolume      float64 `json:"total_volume"  bson:"total_volume"`
	BtcD             string  `json:"btc_d"  bson:"btc_d"`
	EthD             string  `json:"eth_d"  bson:"eth_d"`
	McapChange       string  `json:"mcap_change"  bson:"mcap_change"`
	VolumeChange     string  `json:"volume_change"  bson:"volume_change"`
	AvgChangePercent string  `json:"avg_change_percent"   bson:"avg_change_percent"`
	VolumeAth        int64   `json:"volume_ath" bson:"volume_ath"`
	McapAth          float64 `json:"mcap_ath " bson:"mcap_ath"`
}
