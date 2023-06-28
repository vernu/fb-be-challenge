package main

import (
	"be-assignment-fireb/dbclient"
	"log"

	"be-assignment-fireb/handlers"
	"be-assignment-fireb/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {

	godotenv.Load()

	err := dbclient.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	router := gin.Default()

	initializeCronJob()

	router.GET("/rates/:cryptocurrency/:fiat", handlers.GetExchangeRate)
	router.GET("/rates/:cryptocurrency", handlers.GetAllExchangeRatesForACrypto)
	router.GET("/rates", handlers.GetAllExchangeRates)
	// router.GET("/rates/history/:cryptocurrency/:fiat", handlers.GetExchangeRateHistory)
	router.GET("/balance/:address", handlers.GetEthBalance)

	router.Run(":8080")
}

func initializeCronJob() {
	c := cron.New()

	c.AddFunc("*/5 * * * *", func() {
		log.Println("Cron Job Running")

		data, err := utils.GetExchangeRates()
		if err != nil {
			log.Println(err)
			return
		}

		err = utils.StoreExchangeRates(data, err)
		if err != nil {
			log.Println(err)
			return
		}
	})

	c.Start()
}
