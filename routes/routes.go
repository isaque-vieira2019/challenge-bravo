package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaque-vieira2019/challenge-bravo/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/updateCurrency", controllers.UpdateMainCurrency)
	r.GET("/exchange/", controllers.ExchangeCurrency)
	r.POST("/exchange/add/", controllers.AddCurrency)
	r.DELETE("/exchange/remove/", controllers.RemoveCurrency)
	r.Run(":8000")
}
