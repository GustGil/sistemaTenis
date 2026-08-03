package crawler

import (
	"fmt"
	"net/url"

	"github.com/gocolly/colly"
)

func VisitOnce(c *colly.Collector, link string) {
	mu.Lock()
	defer mu.Unlock()

	if navLinks[link] {
		return
	}
	if !CanVisit() {
		return
	}
	navLinks[link] = true
	if err := c.Visit(link); err != nil {
		fmt.Println(err)
	}

}

func ProcessLink(c *colly.Collector, link string) {
	link = NormalizeURL(link)

	if link == "" {
		return
	}

	if IsNavLink(link) {
		VisitOnce(c, link)
		return
	}

	if IsProduct(link) {
		if _, existe := shoeUrlList[link]; existe {
			fmt.Println("[CRAWLER][VISIT] Link ja cadastrado no DB", link)
			return
		}
		SaveProduct(link)
		return
	}

}

func NormalizeURL(raw string) string {

	u, err := url.Parse(raw)

	if err != nil {
		return ""
	}

	u.Fragment = ""

	return u.String()

}
