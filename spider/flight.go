package spider

import (
	"book_spider/dao"
	"book_spider/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

//  自动同步航班信息
//  记录船只更新信息

func getDayTime(dayCount int) string {
	nowTime := time.Now()
	getTime := nowTime.AddDate(0, 0, dayCount)
	return getTime.Format("2006-01-02")
}

func initFlight() {
	baseQuery := &model.SailQuery{
		Info: model.SailInfo{
			ShipLineCod:   "HKHA",
			PortCodFrom:   "102",
			CarType:       "1",
			TicketType:    "C",
			DepartureDate: "2021-08-08",
		},
	}

	// baseQuery := {"basicParams":{"app":"ctrip","bigChannel":"ship","smallChannel":"bus_index","operatSystem":"android","bigClientType":"h5","smallClientType":"","clientVersion":"3.3.5"},"from":"海口","to":"徐闻","date":"2021-09-02"}

	for i := 0; i <= 7; i++ {
		dateTimeStr := getDayTime(i)

		baseQuery.Info.DepartureDate = dateTimeStr
		queryJSONStr, err := json.Marshal(baseQuery)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		fmt.Println("当前的请求参数:", string(queryJSONStr))
		loadFlightData(queryJSONStr)
	}
}

func loadFlightData(data []byte) {
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"))
	c.OnResponse(func(resp *colly.Response) {
		var resultData model.SailResponse
		err := json.Unmarshal(resp.Body, &resultData)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Visited Data:", resultData.Data)
		// 存储数据
		if resultData.Data != nil {
			dao.SaveFlights(resultData.Data)
			// 获取船的信息
			for _, flight := range resultData.Data {
				fmt.Printf("加载 %s 船的信息:\n", flight.ShipNam)
				getShipDetail(flight.ShipNam)
			}
		}
	})

	c.OnRequest(func(req *colly.Request) {
		req.Headers.Set("Content-Type", "application/json; charset=utf-8")
		fmt.Println("Request", req)
	})

	c.OnError(func(response *colly.Response, e error) {
		fmt.Println("当前返回的错误数据", string(response.Body))
		fmt.Println(e)
	})
	// 执行请求数据
	c.PostRaw("https://m.ctrip.com/restapi/soa2/15123/getShipList", data)
}
