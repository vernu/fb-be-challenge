package utils

import (
	"be-assignment-fireb/dbclient"
	"be-assignment-fireb/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetExchangeRates() (map[string]map[string]float64, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd,eur,gbp"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get exchange rates, status code: %d", resp.StatusCode)
	}

	var data map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func StoreExchangeRates(data map[string]map[string]float64, err error) error {
	collection := dbclient.Database.Collection("exchange_rates")

	if err != nil {
		return err
	}

	timestamp := primitive.NewDateTimeFromTime(time.Now())

	for crypto, rates := range data {
		for fiat, rate := range rates {
			exchangeRate := &models.ExchangeRate{
				Cryptocurrency: crypto,
				Fiat:           fiat,
				Rate:           rate,
				CreatedAt:      timestamp,
			}
			_, err := collection.InsertOne(context.Background(), exchangeRate)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
