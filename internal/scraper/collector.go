package scraper

import (
	"log"
	"net/http/cookiejar"
	"time"

	"github.com/gocolly/colly"
)

func initCollector(allowedDomains string) *colly.Collector {
	jar, _ := cookiejar.New(nil)

	c := colly.NewCollector(
		colly.AllowedDomains(allowedDomains),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36"),
	)

	c.SetCookieJar(jar)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		RandomDelay: 2 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/122.0.0.0 Safari/537.36")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml")
		r.Headers.Set("Accept-Language", "pt-BR,pt;q=0.9,en-US;q=0.8")
		r.Headers.Set("Referer", "https://www.nike.com.br/")
		r.Headers.Set("Connection", "keep-alive")
		log.Println("[SCRAPER][VISIT]", r.URL.String())
	})

	return c

}
