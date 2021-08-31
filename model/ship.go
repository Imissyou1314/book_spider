package model

// 船的详细信息返回数据体
type ShipResponse struct {
	Status int    `json:"status"`
	Data   []Ship `json:"data"`
}

// 船的照片信息返回体
type ShipPicnoResponse struct {
	Status int         `json:"status"`
	Data   []ShipPhoto `json:"data"`
}

// 船的照片相关信息
type ShipPhoto struct {
	Picno      string `json:"picno"`
	Username   string `json:"username"`
	Phototime  string `json:"phototime"`
	Subtime    string `json:"subtime"`
	Photoplace string `json:"photoplace"`
	Anonymous  string `json:"Anonymous"`
	URL        string `json:"url"`
	Surl       string `json:"surl"`
}

// 船的基本信息
type Ship struct {
	Source       int    `json:"source"`
	Mmsi         string `json:"mmsi"`
	Shipid       string `json:"shipid"`
	Tradetype    int    `json:"tradetype"`
	Type         int    `json:"type"`
	Imo          string `json:"imo"`
	Name         string `json:"name"`
	Matchtype    int    `json:"matchtype"`
	Cnname       string `json:"cnname"`
	Callsign     string `json:"callsign"`
	Length       int    `json:"length"`
	Width        int    `json:"width"`
	Left         int    `json:"left"`
	Trail        int    `json:"trail"`
	Draught      int    `json:"draught"`
	Dest         string `json:"dest"`
	Eta          string `json:"eta"`
	Laststa      int    `json:"laststa"`
	Lon          int    `json:"lon"`
	Lat          int    `json:"lat"`
	Sog          int    `json:"sog"`
	Cog          int    `json:"cog"`
	Hdg          int    `json:"hdg"`
	Rot          int    `json:"rot"`
	Navistatus   int    `json:"navistatus"`
	Lastdyn      int    `json:"lastdyn"`
	Satelliteutc int    `json:"satelliteutc"`
	PictureURL   string
	SPictureURL  string
}
