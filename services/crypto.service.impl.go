package services

import (
	"context"
	"crypto-api/models"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

// CryptoServiceImpl is an implementation of the CryptoService interface.
type CryptoServiceImpl struct {
	collection *mongo.Collection
}

// GetExternalCryptoData implements CryptoService.
func (*CryptoServiceImpl) GetExternalCryptoData(*string) (*models.CryptoPrice, error) {
	panic("unimplemented")
}

// NewCryptoService creates a new instance of CryptoServiceImpl.
func NewCryptoService(db *mongo.Database) *CryptoServiceImpl {
	collection := db.Collection("cryptodata")
	return &CryptoServiceImpl{collection}
}

// CreateData creates a new crypto record.
func (s *CryptoServiceImpl) CreateData(crypto *models.CryptoPrice) error {
	_, err := s.collection.InsertOne(context.TODO(), crypto)
	if err != nil {
		return err
	}
	return nil
}

// GetData retrieves a crypto record by ID.
func (s *CryptoServiceImpl) GetData(id *string) (*models.CryptoPrice, error) {
	filter := bson.M{"_id": id} // Assuming "_id" is the identifier field
	var result models.CryptoPrice
	err := s.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("crypto record not found")
		}
		return nil, err
	}
	return &result, nil
}

// get all record from endpoint
func GetExternalCryptoData(baseURL string) ([]models.CryptoPrice, error) {
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

// GetAll retrieves all crypto records.
func (s *CryptoServiceImpl) GetAll() ([]*models.CryptoPrice, error) {
	findOptions := options.Find()
	var cryptos []*models.CryptoPrice
	cursor, err := s.collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var crypto models.CryptoPrice
		err := cursor.Decode(&crypto)
		if err != nil {
			return nil, err
		}
		cryptos = append(cryptos, &crypto)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return cryptos, nil
}

// UpdateData updates an existing crypto record.
func (s *CryptoServiceImpl) UpdateData(crypto *models.CryptoPrice) error {
	filter := bson.M{"_id": crypto.ID} // Assuming "_id" is the identifier field
	update := bson.M{"$set": crypto}
	_, err := s.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// DeleteData deletes a crypto record by ID.
func (s *CryptoServiceImpl) DeleteData(id *string) error {
	filter := bson.M{"_id": id} // Assuming "_id" is the identifier field
	_, err := s.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
