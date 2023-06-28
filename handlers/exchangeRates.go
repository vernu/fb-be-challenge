package handlers

import (
	"be-assignment-fireb/dbclient"
	"be-assignment-fireb/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetExchangeRate(c *gin.Context) {
	cryptocurrency := c.Param("cryptocurrency")
	fiat := c.Param("fiat")

	collection := dbclient.Database.Collection("exchange_rates")

	filter := bson.M{
		"cryptocurrency": cryptocurrency,
		"fiat":           fiat,
	}

	opts := options.FindOne().SetSort(bson.D{{"createdAt", -1}})

	var exchangeRate models.ExchangeRate
	if err := collection.FindOne(context.Background(), filter, opts).Decode(&exchangeRate); err != nil {
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

	opts := options.Find().SetSort(bson.D{{"createdAt", -1}})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get exchange rates",
		})
		return
	}

	var exchangeRates []models.ExchangeRate
	cursor.All(context.Background(), &exchangeRates)

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

	opts := options.Find().SetSort(bson.D{{"createdAt", -1}})
	cursor, err := collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get exchange rates",
		})
		return
	}
	defer cursor.Close(context.Background())

	var exchangeRates []models.ExchangeRate
	cursor.All(context.Background(), &exchangeRates)

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

func GetExchangeRateHistory(c *gin.Context) {
	cryptocurrency := c.Param("cryptocurrency")
	fiat := c.Param("fiat")

	collection := dbclient.Database.Collection("exchange_rates")

	opts := options.Find().SetSort(bson.D{{"createdAt", 1}})

	filter := bson.M{
		"cryptocurrency": cryptocurrency,
		"fiat":           fiat,
		"createdAt": bson.M{
			"$gte": primitive.NewDateTimeFromTime(time.Now().Add(-24 * time.Hour)),
		},
	}

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get exchange rates",
		})
		return
	}
	defer cursor.Close(context.Background())

	var exchangeRates []models.ExchangeRate
	cursor.All(context.Background(), &exchangeRates)

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
