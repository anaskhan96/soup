# soup
[![Build Status](https://travis-ci.org/anaskhan96/soup.svg?branch=master)](https://travis-ci.org/anaskhan96/soup)
[![GoDoc](https://godoc.org/github.com/anaskhan96/soup?status.svg)](https://godoc.org/github.com/anaskhan96/soup)
[![Go Report Card](https://goreportcard.com/badge/github.com/anaskhan96/soup)](https://goreportcard.com/report/github.com/anaskhan96/soup)

**Web Scraper in Go, similar to BeautifulSoup**

*soup* is a small web scraper package for Go, with its interface highly similar to that of BeautifulSoup.

Functions implemented till now :
```go
func Get(string) (string,error) // Takes the url as an argument, returns HTML string
func HTMLParse(string) interface{} // Takes the HTML string as an argument, returns a pointer to the DOM constructed
func Find([]string) interface{} // Element tag,(attribute key-value pair) as argument, pointer to first occurence returned
func FindAll([]string) []struct{} // Same as Find(), but pointers to all occurrences returned
func FindNextSibling() interface{} // Pointer to the next sibling of the Element in the DOM returned
func FindPrevSibling() interface{} // Pointer to the previous sibling of the Element in the DOM returned
func Attrs() map[string]string // Map returned with all the attributes of the Element as lookup to their respective values
func Tag() string // Tag name of the Element returned
func Text() string // Full text inside a non-nested tag returned
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
