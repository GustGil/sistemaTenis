package scraper

import (
	"log"
	"sistemaTenis/internal/data"
	"sistemaTenis/internal/product"

	"github.com/gocolly/colly"
)

func ScrapeProduct(c *colly.Collector, link string) (*product.Tenis, error) {
	shoe := product.Tenis{
		Url: link,
	}

	c.OnHTML(`[data-testid="currentPrice-container"]`, func(e *colly.HTMLElement) {
		shoe.Price = e.Text
	})

	c.OnHTML(".product-title", func(e *colly.HTMLElement) {
		shoe.Name = e.Text
	})

	c.OnHTML(`[data-testid="product_subtitle"]`, func(e *colly.HTMLElement) {
		category := data.CleanCategory(e.Text)
		if category == "" {
			log.Println("[SCRAPER][VISIT] erro ao encontrar a categoria, (category == '') ")
			return
		}
		shoe.Category = category
	})

	shoe.Status = "pending"

	err := c.Visit(link)
	if err != nil {
		return nil, err
	}
	c.Wait()

	return &shoe, nil
}
