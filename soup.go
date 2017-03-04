/* soup package implements a simple web scraper for Go,
keeping it as similar as possible to BeautifulSoup
*/

package soup

import (
	"github.com/anaskhan96/soup/fetch"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Node interface {
	Find(args ...string) Node
	NodeValue() string
	Attrs() map[string]string
	Text() string
	FindAll(args ...string) []Root
	FindNextSibling() Node
	FindPrevSibling() Node
	FindNextElementSibling() Node
	FindPrevElementSibling() Node
}

type Root struct {
	Pointer *html.Node
}

// Returns the HTML returned by the url in string
func Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	s := string(bytes)
	return s, nil
}

// Parses the HTML returning a start pointer to the DOM
func HTMLParse(s string) Node {
	r, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(s, "<!") {
		return Root{r.FirstChild.NextSibling}
	}
	return Root{r}
}

// Finds the first occurrence of the given tag name,
// with or without attribute key and value specified,
// and returns a struct with a pointer to it
func (r Root) Find(args ...string) Node {
	temp, ok := fetch.FindOnce(r.Pointer, args, false)
	if ok == false {
		log.Fatal("Element ", args[0], " with attributes ", args[1:], " not found")
	}
	return Root{temp}
}

// Finds all occurrences of the given tag name,
// with or without key and value specified,
// and returns an array of structs, each having
// the respective pointers
func (r Root) FindAll(args ...string) []Root {
	temp := fetch.FindAllofem(r.Pointer, args)
	if len(temp) == 0 {
		return nil
	}
	pointers := make([]Root, 0, 10)
	for i := 0; i < len(temp); i++ {
		pointers = append(pointers, Root{temp[i]})
	}
	return pointers
}

func (r Root) FindNextSibling() Node {
	nextSibling := r.Pointer.NextSibling
	if nextSibling == nil {
		log.Fatal("No next sibling found")
	}
	return Root{nextSibling}
}

func (r Root) FindPrevSibling() Node {
	prevSibling := r.Pointer.PrevSibling
	if prevSibling == nil {
		log.Fatal("No previous sibling found")
	}
	return Root{prevSibling}
}

// Finds the next element sibling of the pointer in the DOM
// returning a struct with a pointer to it
func (r Root) FindNextElementSibling() Node {
	nextSibling := r.Pointer.NextSibling
	if nextSibling == nil {
		log.Fatal("No next sibling found")
	}
	if nextSibling.Type == html.ElementNode {
		return Root{nextSibling}
	} else {
		p := Root{nextSibling}
		return p.FindNextElementSibling()
	}
}

// Finds the previous element sibling of the pointer in the DOM
// returning a struct with a pointer to it
func (r Root) FindPrevElementSibling() Node {
	prevSibling := r.Pointer.PrevSibling
	if prevSibling == nil {
		log.Fatal("No previous sibling found")
	}
	if prevSibling.Type == html.ElementNode {
		return Root{prevSibling}
	} else {
		p := Root{prevSibling}
		return p.FindPrevElementSibling()
	}
}

// Returns the nodeValue of the element
func (r Root) NodeValue() string {
	return r.Pointer.Data
}

// Returns an array containing key and values of all attributes
func (r Root) Attrs() map[string]string {
	if len(r.Pointer.Attr) == 0 {
		return nil
	}
	return fetch.GetKeyValue(r.Pointer.Attr)
}

// Returns the string inside a non-nested element
func (r Root) Text() string {
	k := r.Pointer.FirstChild
	if k.Type == html.TextNode {
		return k.Data
	}
	return ""
}

/* Prettify() function to be looked at later
func Prettify(r io.ReadCloser) (string, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	str := string(bytes)
	var mod string
	for i, c := range str {
		if c == '>' || c == ';' || str[i+1] == '<' {
			mod += string(c)
			mod += "\n"
			continue
		}
		mod += string(c)
	}
	return mod, nil
}
*/
