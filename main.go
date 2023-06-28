package main

import (
	"be-assignment-fireb/dbclient"
	"log"

	"be-assignment-fireb/handlers"
	"be-assignment-fireb/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	err := dbclient.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	router := gin.Default()

	router.GET("/fetch", func(c *gin.Context) {
		data, err := utils.GetExchangeRates()
		err = utils.StoreExchangeRates(data, err)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "data saved successfully",
			"data":    data,
		})
	})
	router.GET("/rates/:cryptocurrency/:fiat", handlers.GetExchangeRate)
	router.GET("/rates/:cryptocurrency", handlers.GetAllExchangeRatesForACrypto)
	router.GET("/rates", handlers.GetAllExchangeRates)
	// router.GET("/rates/history/:cryptocurrency/:fiat", handlers.GetExchangeRateHistory)
	router.GET("/balance/:address", handlers.GetEthBalance)

	router.Run(":8080")
}
