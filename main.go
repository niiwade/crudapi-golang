package main

import (
	"crypto-api/cryptoapi"
	"io"
	"log"
	"os"
)

func main() {

	rosterFile, err := os.OpenFile("crypto-data.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening the file crypto-data.txt: %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	cryptodata, err := cryptoapi.GetAllCryptoData()
	if err != nil {
		log.Fatalf("error while getting all crypto data: %v", err)
	}

	for _, data := range cryptodata {
		log.Println("---------------------")
		log.Printf("CoinsCount: %d\n", data.CoinsCount)
		log.Printf("ActiveMarkets: %d\n", data.ActiveMarkets)
		log.Printf("TotalMcap: %f\n", data.TotalMcap)
		log.Printf("TotalVolume : %f\n", data.TotalVolume)
		log.Printf("BtcD: %s\n", data.BtcD)
		log.Printf("EthD: %s\n", data.EthD)
		log.Printf("McapChange : %s\n", data.McapChange)
		log.Printf("VolumeChange: %s\n", data.VolumeChange)
		log.Printf("AvgChangePercent: %s\n", data.AvgChangePercent)
		log.Printf("VolumeAth : %d\n", data.VolumeAth)
		log.Printf("McapAth  : %f\n", data.McapAth)
		log.Println("---------------------")
	}

}
