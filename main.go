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
		// Add more fields as needed
		log.Println("---------------------")
	}

}
