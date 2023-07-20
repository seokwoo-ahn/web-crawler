package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr > td > em.time", func(i int, elem *colly.HTMLElement) {
			if i%2 == 0 {
				fmt.Print(elem.Text + " ")
			} else {
				fmt.Println(elem.Text)
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Post(
		"https://etk.srail.kr/hpg/hra/01/selectScheduleList.do?pageId=TK0101010000",
		map[string]string{
			"dptRsStnCd":      "0551",
			"arvRsStnCd":      "0020",
			"stlbTrnClsfCd":   "05",
			"psgNum":          "1",
			"seatAttCd":       "015",
			"isRequest":       "Y",
			"dptRsStnCdNm":    "수서",
			"arvRsStnCdNm":    "부산",
			"dptDt":           "20230720",
			"dptTm":           "153500",
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
		},
	)
	// c.Visit("https://go-colly.org/")

	// c.OnError()
}
