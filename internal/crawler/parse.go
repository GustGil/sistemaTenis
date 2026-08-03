package crawler

import "strings"

func IsProduct(link string) bool {
	link = strings.ToLower(link)

	if strings.Contains(link, "/t/") && strings.Contains(link, "shoe") {
		return true
	}

	return false
}

func IsNavLink(link string) bool {

	link = strings.ToLower(link)

	if strings.Contains(link, "/w/") {
		return true
	}

	return false
}

func CanVisit() bool {
	mu2.Lock()
	defer mu2.Unlock()

	if navCount >= 300 { // limite total de páginas
		return false
	}

	navCount++
	return true
}
