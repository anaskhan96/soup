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
	Tag() string
	Attrs() []html.Attribute
	Text() string
	FindAll(args ...string) []Root
	FindNextSibling() Node
	FindPrevSibling() Node
}

type Root struct {
	Pointer *html.Node
}

// Returns the HTML returned by the url in string
func Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "<>", err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "<>", err
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
	temp, ok, _ := fetch.FindOnce(r.Pointer, args, false)
	if ok == false {
		return nil
	}
	return Root{temp}
}

// Finds all occurrences of the given tag name,
// with or without key and value specified,
// and returns an array of structs, each having
// the respective pointers
func (r Root) FindAll(args ...string) []Root {
	fetch.Set()
	temp, _, _:= fetch.FindAllofem(r.Pointer, args, false)
	if len(temp) == 0 {
		return nil
	}
	pointers := make([]Root, 0, 10)
	for i := 0; i < len(temp); i++ {
		pointers = append(pointers, Root{temp[i]})
	}
	return pointers
}

// Finds the next sibling of the pointer in the DOM
// returning a struct with a pointer to it
func (r Root) FindNextSibling() Node {
	return Root{r.Pointer.NextSibling.NextSibling}
}

// Finds the previous sibling of the pointer in the DOM
// returning a struct with a pointer to it
func (r Root) FindPrevSibling() Node {
	return Root{r.Pointer.PrevSibling.PrevSibling}
}

// Returns the Tag name of the element
func (r Root) Tag() string {
	return r.Pointer.Data
}

// Returns an array containing key and values of all attributes
func (r Root) Attrs() []html.Attribute {
	return r.Pointer.Attr
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
