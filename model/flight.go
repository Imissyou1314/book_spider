package model

type SailQuery struct {
	Info SailInfo `json:"info"`
}
type SailInfo struct {
	ShipLineCod   string `json:"ShipLineCod"`
	PortCodFrom   string `json:"PortCod_from"`
	CarType       string `json:"CarType"`
	TicketType    string `json:"TicketType"`
	DepartureDate string `json:"DepartureDate"`
}

type SailResponse struct {
	Data     []Flight `json:"Data"`
	Code     int      `json:"Code"`
	Message  string   `json:"Message"`
	DateTime string   `json:"DateTime"`
}

type Flight struct {
	ID            string `json:"ID"`
	SaleVoyageCod string `json:"SaleVoyageCod"`
	SailDate      string `json:"SailDate"`
	SailTime      string `json:"SailTime"`
	ShipLineCod   string `json:"ShipLineCod"`
	ShipCod       string `json:"ShipCod"`
	ShipNam       string `json:"ShipNam"`
	PortCod       string `json:"PortCod"`
	ToPortCod     string `json:"ToPortCod"`
	PortCodNam    string `json:"PortCodNam"`
	ToPortCodNam  string `json:"ToPortCodNam"`
	SellStatus    string `json:"SellStatus"`
	TicketNum     string `json:"TicketNum"`
}
