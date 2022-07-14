package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateMainCurrency(c *gin.Context) {
	resp, err := http.Get("https://economia.awesomeapi.com.br/last/BRL-USD,EUR-USD,BTC-USD,ETH-USD")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	type Temp struct {
		code  string `JSON:"code"`
		value string `JSON:"bid"`
	}

	respData, err := ioutil.ReadAll(resp.Body)

	var temp Temp
	json.Unmarshal(respData, &temp)
	c.JSON(http.StatusOK, temp)

}

func ExchangeCurrency(c *gin.Context) {

}

func AddCurrency(c *gin.Context) {

}

func RemoveCurrency(c *gin.Context) {

}
