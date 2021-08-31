package router

import (
	"book_spider/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getShipDetail(c *gin.Context) {
	shipName := c.PostForm("shipName")
	if shipName != "" {
		var ship models.Ship
		dao.GetShipByName(shipName, &ship)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ship})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "", "message": "当前请求参数错误"})
	}
}
