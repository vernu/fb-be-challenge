package handlers

import (
	"be-assignment-fireb/utils"

	"github.com/gin-gonic/gin"
)

func GetEthBalance(c *gin.Context) {
	address := c.Param("address")

	balance, err := utils.GetEthBalance(address)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"balance": balance,
	})
}
