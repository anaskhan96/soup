# soup
**Web Scraper in Go, similar to BeautifulSoup in Python**

*soup* is a small web scraper package for Go, with its interface highly similar to that of BeautifulSoup in Python.

Functions implemented till now :
```go
func Get(string) (string,error)
func HTMLParse(string) interface{}
func Find([]string) interface{}
func FindAll([]string) []struct{}
func FindNextSibling() []interface{}
func FindPrevSibling() []interface{}
func Attrs() map[string]string
func Tag() string
func Text() string
```

## Installation
Install the package using the command
```bash
go get github.com/anaskhan96/soup
```

## Example
An example code is given below to scrape the "Comics I Enjoy" part (text and its links) from [xkcd](https://xkcd.com).

[More Examples](https://github.com/anaskhan96/soup/tree/master/examples)
```go
package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func main() {
	resp, err := soup.Get("https://xkcd.com")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "id", "comicLinks").FindAll("a")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	}
}
```
