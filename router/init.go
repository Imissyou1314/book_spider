package router

import (
	"book_spider/config"

	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()

	apiRouter := router.Group("/api/book")

	{
		apiRouter.POST("/flights", loadFlights)
		apiRouter.POST("/shipDetail", getShipDetail)
	}
	port := config.GetConfig("port", "3388")
	router.Run(":" + port)
}
