package main

import (
	"context"
	"crypto-api/controllers"
	"crypto-api/cryptoapi"
	"crypto-api/routes"
	"crypto-api/services"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// Initialize the Gin router
	router := gin.Default()

	// Initialize the MongoDB client and connect to your database
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB server: %v", err)
	}

// Access your MongoDB database
db := client.Database("Crytpodb") // Replace with your database name

// Initialize the CryptoService with the MongoDB database
cryptoService := services.NewCryptoService(db)

// Initialize the CryptoController with the CryptoService
cryptoController := controllers.NewCryptoController(cryptoService)


	// CORS middleware,
	// router.Use(cors.Default())

	// Set up API routes
	routes.SetupRouter(router, cryptoController)

	// Start the Gin server on a specific port (e.g., :8080)
	log.Printf("Starting server on port 8080...")
	router.Run(":8080")

}
