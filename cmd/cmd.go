package main

import (
	"web-crawler/crawler"
)

var Crawler *crawler.Crawler
var StationCodes map[string]string

func Init() {
	Crawler, _ = crawler.NewCrawler()
}

func main() {
	Init()
	Crawler.Run()
}
