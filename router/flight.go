package router

import (
	"book_spider/dao"
	"book_spider/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func loadFlights(c *gin.Context) {
	dateStr := c.PostForm("date")
	fmt.Println("当前请求参数:", dateStr)
	if dateStr != "" {
		var resultDate []model.Flight
		dao.LoadFlightsByDate(dateStr, &resultDate)
		fmt.Println(resultDate)
		if len(resultDate) > 0 {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": resultDate})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "", "message": "当天无数据"})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "", "message": "请求参数错误"})
	}
}
