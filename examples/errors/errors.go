// Errors happen. This example shows how to detect and handle some of them.

package main

import (
	"fmt"
	"log"

	"github.com/anaskhan96/soup"
)

func main() {
	_, err := soup.Get("this url isn't real!")
	if err != nil && err.(soup.Error).Type == soup.ErrInGetRequest {
		// Handle as required!
	}

	url := fmt.Sprintf("https://xkcd.com/50")
	xkcd, err := soup.Get(url)
	if err != nil {
		// Handle it
	}
	xkcdSoup := soup.HTMLParse(xkcd)
	links := xkcdSoup.Find("div", "id", "linkz")
	if links.Error != nil && links.Error.(soup.Error).Type == soup.ErrElementNotFound {
		log.Printf("Element not found: %v", links.Error)
	}
	// These error types were introduced in version 1.2.0, but just checking for err still works:
	links = xkcdSoup.Find("div", "id", "links2")
	if links.Error != nil {
		log.Printf("Something happened: %s", links.Error)
	}
}
