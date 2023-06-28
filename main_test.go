package main

import (
	"be-assignment-fireb/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	godotenv.Load()
	gin.SetMode(gin.TestMode)
	m.Run()

}

func TestGetEthBalanceWithValidAddress(t *testing.T) {
	r := gin.Default()

	r.GET("/balance/:address", handlers.GetEthBalance)
	req, _ := http.NewRequest("GET", "/balance/0x742d35Cc6634C0532925a3b844Bc454e4438f44e", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code)

}

func TestGetEthBalanceWithInvalidAddress(t *testing.T) {
	r := gin.Default()

	r.GET("/balance/:address", handlers.GetEthBalance)
	req, _ := http.NewRequest("GET", "/balance/invalidAddress", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, 400, res.Code)
	assert.JSONEq(t, `{"error": "invalid ethereum address"}`, res.Body.String())
}
