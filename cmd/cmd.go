package main

import (
	"web-crawler/crawler"
)

var Crawler *crawler.Crawler

func Init() {
	Crawler, _ = crawler.NewCrawler()
	Crawler.SetCrawler()
}

func main() {
	Init()
	Crawler.Run()
}
