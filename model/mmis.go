package model

type MMISResponse struct {
	Status int `json:"status"`
	Ship   []struct {
		M   int    `json:"m"` // mmis 编码
		N   string `json:"n"` // 船名称
		I   int    `json:"i"`
		C   string `json:"c"`
		T   int    `json:"t"`
		QTY string `json:"QTY"`
	} `json:"ship"`
	Port []interface{} `json:"port"`
}
