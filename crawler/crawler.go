package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Crawler struct {
	collector *colly.Collector
	reqBody   map[string]string
	url       string
}

func NewCrawler() (crawler *Crawler, err error) {
	collector := colly.NewCollector(
		colly.AllowedDomains("https://etk.srail.kr", "etk.srail.kr"),
	)

	reqBody := map[string]string{
		"dptRsStnCd":      "0551",
		"arvRsStnCd":      "0020",
		"stlbTrnClsfCd":   "05",
		"psgNum":          "1",
		"seatAttCd":       "015",
		"isRequest":       "Y",
		"prvTms":          "133000",
		"dptRsStnCdNm":    "수서",
		"arvRsStnCdNm":    "부산",
		"dptDt":           "20230725",
		"dptTm":           "133000",
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

	err = nil
	return
}

func (c *Crawler) SetCrawler() {
	c.collector.OnHTML("table", func(e *colly.HTMLElement) {
		count := 0
		e.ForEach("tr > td > em.time", func(i int, elem *colly.HTMLElement) {
			if i%2 == 0 {
				fmt.Print(elem.Text + " ")
			} else {
				fmt.Println(elem.Text)
			}
			count++
		})
		// if count != 20 {
		// 	return
		// } else {
		// 	fmt.Println(e.Response.Ctx)
		// 	e.Response.Ctx.ForEach(func(k string, v interface{}) interface{} {
		// 		fmt.Println(k)
		// 		fmt.Println(v)
		// 		return v
		// 	})
		// }
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
