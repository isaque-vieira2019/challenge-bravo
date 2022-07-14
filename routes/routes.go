package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaque-vieira2019/challenge-bravo/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/exchange/", controllers.ExchangeCurrency)
	r.GET("/exchange/add/", controllers.AddCurrency)
	r.GET("/exchange/remove/", controllers.RemoveCurrency)
}
