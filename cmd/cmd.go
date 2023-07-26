package main

import (
	"flag"
	"web-crawler/config"
	"web-crawler/crawler"
)

var srtCrawler *crawler.Crawler
var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func Init() {
	config := config.NewConfig(*configFlag)
	srtCrawler, _ = crawler.NewCrawler(config)
}

func main() {
	Init()
	srtCrawler.Run()
}
