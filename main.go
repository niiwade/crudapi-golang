package main

import (
	"io"
	"log"
	"metro-api/metroapi"
	"os"
)

func main() {

	rosterFile, err := os.OpenFile("weather.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening the file weather.txt: %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	cryptodata, err := metroapi.GetAllCryptoData()
	if err != nil {
		log.Fatalf("error while getting all weather: %v", err)
	}

	for _, data := range cryptodata {
		log.Println("---------------------")
		log.Printf("CoinsCount: %d\n", data.CoinsCount)
		log.Printf("ActiveMarkets: %d\n", data.ActiveMarkets)
		// Add more fields as needed
		log.Println("---------------------")
	}

}
