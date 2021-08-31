package dao

import (
	"book_spider/model"
	"fmt"
)

type FlightDao struct{}

func SaveFlights(flights []model.Flight) {
	db.Create(flights)
}

func LoadFlightsByDate(dateStr string, flights *[]model.Flight) {
	fmt.Println("当前请求参数:", dateStr)
	db.Where("sail_date = ?", dateStr).Find(flights)
}
