package crawler

import (
	"sistemaTenis/internal/product"
	"sync"

	"go.mongodb.org/mongo-driver/v2/bson"
)

var (
	navLinks     = make(map[string]bool)
	productLinks = make(map[string]bool)
	mu           sync.Mutex
	navCount     int
	mu2          sync.Mutex
	shoeUrlList  map[any]struct{}
)

func Init(startUrl string) map[string]bool {

	c := InitCollector("www.nike.com")

	shoeUrlList = product.GetShoesByParam(&bson.M{
		"url": 1,
		"_id": 0,
	})

	VisitOnce(c, startUrl)
	c.Wait()

	return productLinks
}
