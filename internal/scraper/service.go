package scraper

import (
	"fmt"
	"log"
	"sistemaTenis/internal/product"
)

func Init(productLink map[string]bool, allowedDomains string) {

	c := initCollector(allowedDomains)
	log.Println("[SCRAPER] DADOS EXTRAIDOS:")

	for link := range productLink {
		shoe, err := ScrapeProduct(c, link)
		if err != nil {
			continue
		}
		fmt.Printf("nome: %s - %s - %s - %s \n", shoe.Name, shoe.Price, shoe.Category, shoe.Url)
		product.SaveBasic(shoe)
	}
}
