package spider

import (
	"book_spider/dao"
	"book_spider/model"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

//  自动同步航班信息
//  记录船只更新信息

func getShipDetail(shipName string) {
	// config Data
	var mmsiStr string
	// 获取ship mmsi
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"))
	var mmsiData model.MMISResponse
	c.Cookies("http://www.shipxy.com")

	// 获取ship 详情的请求
	c2 := c.Clone()
	var shipData model.ShipResponse
	c2.Cookies("http://www.shipxy.com")

	// 获取图片的请求
	c3 := c.Clone()
	var shipPhoto model.ShipPicnoResponse
	c3.Cookies("http://www.shipxy.com")

	c3.OnResponse(func(resp *colly.Response) {
		fmt.Println("LOAD Ship Photo ===========================>", string(resp.Body))
		err := json.Unmarshal(resp.Body, &shipPhoto)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("LOAD Ship Photo:", shipPhoto.Data)
		if len(shipPhoto.Data) > 0 {
			saveShipDate(shipData.Data[0], shipPhoto.Data[0])
		} else {
			saveShipDate(shipData.Data[0], model.ShipPhoto{})
		}
	})

	c2.OnRequest(func(req *colly.Request) {
		req.Headers.Set("content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		req.Headers.Set("accept", "application/json, text/javascript, */*; q=0.01")
		req.Headers.Set("cookie", "FD857C2AF68165D4=grjFYD3TGLix0Zlj3n234W6b3IZ4e3bS1C5weR5mSPtpHo6qDp2eXW/DJnjxUu+FIN430HIHWGg=; _elane_maptype=MT_GOOGLE; Hm_lvt_adc1d4b64be85a31d37dd5e88526cc47=1609732197; tc_TC=; _elane_shipfilter_type=%u8D27%u8239%2C%u96C6%u88C5%u7BB1%u8239%2C%u6CB9%u8F6E%2C%u5F15%u822A%u8239%2C%u62D6%u8F6E%2C%u62D6%u5F15%2C%u6E14%u8239%2C%u6355%u635E%2C%u5BA2%u8239%2C%u641C%u6551%u8239%2C%u6E2F%u53E3%u4F9B%u5E94%u8239%2C%u88C5%u6709%u9632%u6C61%u88C5%u7F6E%u548C%u8BBE%u5907%u7684%u8239%u8236%2C%u6267%u6CD5%u8247%2C%u5907%u7528-%u7528%u4E8E%u5F53%u5730%u8239%u8236%u7684%u4EFB%u52A1%u5206%u914D%2C%u5907%u7528-%u7528%u4E8E%u5F53%u5730%u8239%u8236%u7684%u4EFB%u52A1%u5206%u914D%2C%u533B%u7597%u8239%2C%u7B26%u540818%u53F7%u51B3%u8BAE%28Mob-83%29%u7684%u8239%u8236%2C%u62D6%u5F15%u5E76%u4E14%u8239%u957F%3E200m%u6216%u8239%u5BBD%3E25m%2C%u758F%u6D5A%u6216%u6C34%u4E0B%u4F5C%u4E1A%2C%u6F5C%u6C34%u4F5C%u4E1A%2C%u53C2%u4E0E%u519B%u4E8B%u884C%u52A8%2C%u5E06%u8239%u822A%u884C%2C%u5A31%u4E50%u8239%2C%u5730%u6548%u5E94%u8239%2C%u9AD8%u901F%u8239%2C%u5176%u4ED6%u7C7B%u578B%u7684%u8239%u8236%2C%u5176%u4ED6; _elane_shipfilter_length=0%2C40%2C41%2C80%2C81%2C120%2C121%2C160%2C161%2C240%2C241%2C320%2C321%2C9999; _elane_shipfilter_sog=0%2C1; _filter_flag=-1; _elane_shipfilter_one=2; _elane_shipfilter_country=0%2C1%2C2; _elane_shipfilter_olength=; tc_QX=; shipxy_v3_history_serch=s%u2606HAI%20KOU%2016%20HAO%u2606413525630%u260660%u2606MMSI%uFF1A413525630; ASP.NET_SessionId=g0dc2b4s25bm5ym1d33kxtxo; Hm_lpvt_adc1d4b64be85a31d37dd5e88526cc47=1609741906; SERVERID=8dac9e937b8701fba4cb1394cffade3e|1609742245|1609739001")
	})

	c2.OnResponse(func(resp *colly.Response) {
		fmt.Println("LOAD Ship Info ===========================>", string(resp.Body))
		err := json.Unmarshal(resp.Body, &shipData)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("LOAD Ship Data:", shipData.Data)
		photoURL := "http://shippic.shipxy.com/getshippiclisttop.ashx?top=1&mmsi=" + mmsiStr
		c3.Post(photoURL, nil)
	})

	// 请求获取 船的mmis 编码
	c.OnResponse(func(resp *colly.Response) {
		fmt.Println("LOAD MMIS ===========================>", string(resp.Body))
		err := json.Unmarshal(resp.Body, &mmsiData)
		if err != nil {
			fmt.Println(err.Error())
		}
		mmsiStr = strconv.Itoa(mmsiData.Ship[0].M)
		oldShipData := dao.GetShipByMmis(mmsiStr)
		fmt.Println("load dao ship detail:", oldShipData)
		if oldShipData == nil || oldShipData.Mmsi == "" {
			query := map[string]string{"mmsi": mmsiStr}
			fmt.Println("load ship detail:", query)
			c2.Post("http://www.shipxy.com/ship/GetShip", map[string]string{"mmsi": mmsiStr})
		}
	})
	c.OnError(func(response *colly.Response, e error) {
		fmt.Println("当前返回的错误数据", string(response.Body))
		fmt.Println(e)
	})
	showDateilURL := "http://searchv3.shipxy.com/shipdata/search3.ashx?f=auto&kw=" + shipName
	c.Post(showDateilURL, nil)
}

func saveShipDate(shipData model.Ship, shipPhoto model.ShipPhoto) {
	if shipPhoto != (model.ShipPhoto{}) {
		shipData.PictureURL = shipPhoto.URL
		shipData.SPictureURL = shipPhoto.Surl
	}
	dao.SaveShip(shipData)
}
