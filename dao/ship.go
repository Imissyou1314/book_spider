package dao

import (
	"book_spider/model"
)

func GetShipByName(shipName string) *model.Ship {
	var ship = &model.Ship{}
	db.Where("name = ?", shipName).First(ship)
	return ship
}

func GetShipByMmis(mmsi string) *model.Ship {
	var ship = &model.Ship{}
	db.Where("mmsi = ?", mmsi).First(ship)
	return ship
}

func SaveShips(ships []model.Ship) {
	db.Create(ships)
}

func SaveShip(ship model.Ship) {
	db.Create(ship)
}
