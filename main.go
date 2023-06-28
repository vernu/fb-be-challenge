package main

import (
	"be-assignment-fireb/dbclient"
	"log"

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

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola",
		})
	})
	router.Run(":8080")
}
