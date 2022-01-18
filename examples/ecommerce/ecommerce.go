// E-Commerce Example. try to extract products description and handle error when inconsistent contents are not found.
package main

import (
	"log"
	"os"

	"github.com/anaskhan96/soup"
)

func main() {
	resp, err := soup.Get("https://webscraper.io/test-sites/e-commerce/allinone")
	if err != nil {
		// handle as your wish
	}
	doc := soup.HTMLParse(resp)
	wrapper := doc.Find("div", "class", "wrapper")
	if wrapper.Error != nil && wrapper.Error.(soup.Error).Type == soup.ErrElementNotFound { // assert error as soup.Error type
		log.Printf("Wrapper element not found: %v\n", wrapper.Error)
		os.Exit(1) // terminate process
	}
	container := wrapper.FindStrict("div", "class", "container test-site") // use FindStrict if element has more than 1 class
	if container.Error != nil {
		log.Printf("Container element not found: %s\n" + container.Error.Error()) // print original soup.Error msg
		os.Exit(1)
	}
	mainRow := container.Find("div", "class", "row")
	mainCol := mainRow.Find("div", "class", "col-md-9").Find("div", "class", "row")
	products := mainCol.FindAllStrict("div", "class", "col-sm-4 col-lg-4 col-md-4") // same as FindStrict but return []soup.Root

	for i, product := range products {
		caption := product.Find("div", "class", "caption")
		if caption.Error != nil {
			log.Println(caption.Error.Error())
			os.Exit(1)
		}
		description := caption.Find("p", "class", "description")
		log.Printf("Product #%d \n description: %s\n", i, description.Text())
	}
}
