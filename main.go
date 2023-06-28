package main

import (
	"be-assignment-fireb/dbclient"
	"log"

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
	router.Run(":8080")
}
