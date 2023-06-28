package utils

import (
	"be-assignment-fireb/dbclient"
	"be-assignment-fireb/models"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
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

func GetEthBalance(address string) (*big.Float, error) {

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + os.Getenv("INFURA_KEY"))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if !common.IsHexAddress(address) {
		return nil, fmt.Errorf("invalid ethereum address")
	}

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		panic(err)
	}

	balanceInEther := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(params.Ether))

	return balanceInEther, nil

}
