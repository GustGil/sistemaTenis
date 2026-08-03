package crawler

import (
	"log"
)

func SaveProduct(link string) {
	mu.Lock()
	defer mu.Unlock()

	if productLinks[link] {
		return
	}

	productLinks[link] = true
	log.Println("[PRODUCT FOUND]", link)
}
