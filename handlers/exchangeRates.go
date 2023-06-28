package handlers

import (
	"be-assignment-fireb/dbclient"
	"be-assignment-fireb/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetExchangeRate(c *gin.Context) {
	cryptocurrency := c.Param("cryptocurrency")
	fiat := c.Param("fiat")

	collection := dbclient.Database.Collection("exchange_rates")

	filter := bson.M{
		"cryptocurrency": cryptocurrency,
		"fiat":           fiat,
	}

	var exchangeRate models.ExchangeRate
	if err := collection.FindOne(context.Background(), filter).Decode(&exchangeRate); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get exchange rate",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cryptocurrency": cryptocurrency,
		"fiat":           fiat,
		"rate":           exchangeRate.Rate,
	})
}

func GetAllExchangeRatesForACrypto(c *gin.Context) {
	cryptocurrency := c.Param("cryptocurrency")

	collection := dbclient.Database.Collection("exchange_rates")

	filter := bson.M{
		"cryptocurrency": cryptocurrency,
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get exchange rates",
		})
		return
	}

	var exchangeRates []models.ExchangeRate
	if err := cursor.All(context.Background(), &exchangeRates); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get exchange rates ",
		})
		return
	}

	result := make(map[string]float64)
	for _, exchangeRate := range exchangeRates {
		result[exchangeRate.Fiat] = exchangeRate.Rate
	}

	c.JSON(http.StatusOK, gin.H{
		"cryptocurrency": cryptocurrency,
		"rates":          result,
	})
}

func GetAllExchangeRates(c *gin.Context) {

	collection := dbclient.Database.Collection("exchange_rates")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get exchange rates",
		})
		return
	}
	defer cursor.Close(context.Background())

	var exchangeRates []models.ExchangeRate
	if err := cursor.All(context.Background(), &exchangeRates); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get exchange rates",
		})
		return
	}

	result := make(map[string]map[string]float64)
	for _, exchangeRate := range exchangeRates {
		if _, ok := result[exchangeRate.Cryptocurrency]; !ok {
			result[exchangeRate.Cryptocurrency] = make(map[string]float64)
		}
		result[exchangeRate.Cryptocurrency][exchangeRate.Fiat] = exchangeRate.Rate
	}

	c.JSON(http.StatusOK, gin.H{
		"rates": result,
	})
}
