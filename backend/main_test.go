package main

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetSummonerIDPass(t *testing.T) {
	LoadAuth()

	t.Log("Testing GetSummonerID Function passes when expected")
	responseData, err := getSummonerID("Blubbystr")
	assert.Equal(t, err, nil)
	assert.NotEmpty(t, responseData)
}

func TestGetSummonerIDFail(t *testing.T) {
	LoadAuth()

	t.Log("Testing GetSummonerID Function fails when expected")
	_, err := getSummonerID("askldnj02943hqa408d3da0")
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "404 - Summoner Not Found")
}
