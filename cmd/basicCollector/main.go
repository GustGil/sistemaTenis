package main

import (
	"sistemaTenis/internal/crawler"
	"sistemaTenis/internal/scraper"
	"sistemaTenis/repository"
)

func main() {
	repository.Connect()
	List := crawler.Init("https://www.nike.com/")
	scraper.Init(List, "www.nike.com")
	repository.Disconnect()
}
