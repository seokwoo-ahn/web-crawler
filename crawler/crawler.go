package crawler

import (
	"fmt"
	"web-crawler/config"

	"github.com/gocolly/colly"
)

type Crawler struct {
	collector *colly.Collector
	reqBody   map[string]string
	url       string
}

func NewCrawler(config *config.Configs) (crawler *Crawler, err error) {
	collector := colly.NewCollector(
		colly.AllowedDomains("https://etk.srail.kr", "etk.srail.kr"),
	)

	fmt.Println(config.StationMap[config.DptStation])
	fmt.Println(config.StationMap[config.ArvStation])
	fmt.Println(config.DptStation)
	fmt.Println(config.ArvStation)
	fmt.Println(config.DptDay)

	reqBody := map[string]string{
		"dptRsStnCd":      config.StationMap[config.DptStation],
		"arvRsStnCd":      config.StationMap[config.ArvStation],
		"stlbTrnClsfCd":   "05",
		"psgNum":          "1",
		"seatAttCd":       "015",
		"isRequest":       "Y",
		"prvTms":          "000000",
		"dptRsStnCdNm":    config.DptStation,
		"arvRsStnCdNm":    config.ArvStation,
		"dptDt":           config.DptDay,
		"dptTm":           "000000",
		"chtnDvCd":        "1",
		"psgInfoPerPrnb1": "1",
		"psgInfoPerPrnb5": "0",
		"psgInfoPerPrnb4": "0",
		"psgInfoPerPrnb2": "0",
		"psgInfoPerPrnb3": "0",
		"locSeatAttCd1":   "000",
		"rqSeatAttCd1":    "015",
		"trnGpCd":         "109",
		"dlayTnumAplFlg":  "Y",
	}

	url := "https://etk.srail.kr/hpg/hra/01/selectScheduleList.do?pageId=TK0101010000"

	crawler = &Crawler{
		collector: collector,
		reqBody:   reqBody,
		url:       url,
	}

	crawler.SetCrawler()

	return
}

func (c *Crawler) SetCrawler() {
	c.collector.OnHTML("table", func(e *colly.HTMLElement) {
		count := 0
		e.ForEach("tr > td > em.time", func(i int, elem1 *colly.HTMLElement) {
			if i%2 == 0 {
				fmt.Print(elem1.Text + " ")
			} else {
				fmt.Println(elem1.Text)
			}
			count++
		})
		if count != 20 {
			return
		} else {
			e.ForEach("tr > td.trnNo > input[name]", func(j int, elem2 *colly.HTMLElement) {
				if elem2.Attr("name") == "dptTm[0]" {
					c.reqBody["prvTms"] = elem2.Attr("value")
					fmt.Println("set prvTms", elem2.Attr("value"))
				}

				if elem2.Attr("name") == "dptTm[9]" {
					c.reqBody["dptTm"] = elem2.Attr("value")
					fmt.Println("set dptTm", elem2.Attr("value"))
				}
			})
			c.collector.Post(c.url, c.reqBody)
		}
	})

	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	if strings.Contains(e.Attr("href"), "selectScheduleList") {
	// 		e.Request.Visit(e.Attr("href"))
	// 	}
	// })

	c.collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
}

func (c *Crawler) Run() {
	c.collector.Post(c.url, c.reqBody)
}
