/* soup package implements a simple web scraper for Go,
keeping it as similar as possible to BeautifulSoup
*/

package soup

import (
	"github.com/anaskhan96/soup/fetch"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"strings"
)

type Node interface {
	Find(args ...string) Root
	Attrs() map[string]string
	Text() string
	FindAll(args ...string) []Root
	FindNextSibling() Root
	FindPrevSibling() Root
	FindNextElementSibling() Root
	FindPrevElementSibling() Root
}

type Root struct {
	Pointer   *html.Node
	NodeValue string
}

// Returns the HTML returned by the url in string
func Get(url string) (string, error) {
	defer fetch.CatchPanic("Get()")
	resp, err := http.Get(url)
	if err != nil {
		panic("Couldn't perform GET request to "+url)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Unable to read the response body")
	}
	s := string(bytes)
	return s, nil
}

// Parses the HTML returning a start pointer to the DOM
func HTMLParse(s string) Root {
	defer fetch.CatchPanic("HTMLParse()")
	r, err := html.Parse(strings.NewReader(s))
	if err != nil {
		panic("Unable to parse the HTML")
	}
	//Navigate to find an html.ElementNode
	for r.Type != html.ElementNode {
		switch r.Type {
		case html.DocumentNode:
			r = r.FirstChild
		case html.DoctypeNode:
			r = r.NextSibling
		case html.CommentNode:
			r = r.NextSibling
		}

	}
	return Root{r, r.Data}
}

// Finds the first occurrence of the given tag name,
// with or without attribute key and value specified,
// and returns a struct with a pointer to it
func (r Root) Find(args ...string) Root {
	defer fetch.CatchPanic("Find()")
	temp, ok := fetch.FindOnce(r.Pointer, args, false)
	if ok == false {
		panic("Element `" + args[0] + "` with attributes `" + strings.Join(args[1:], " ") + "` not found")
	}
	return Root{temp, temp.Data}
}

// Finds all occurrences of the given tag name,
// with or without key and value specified,
// and returns an array of structs, each having
// the respective pointers
func (r Root) FindAll(args ...string) []Root {
	defer fetch.CatchPanic("FindAll()")
	temp := fetch.FindAllofem(r.Pointer, args)
	if len(temp) == 0 {
		panic("No element `" + args[0] + "` with attributes `" + strings.Join(args[1:], " ") + "` not found")
	}
	pointers := make([]Root, 0, 10)
	for i := 0; i < len(temp); i++ {
		pointers = append(pointers, Root{temp[i], temp[i].Data})
	}
	return pointers
}

func (r Root) FindNextSibling() Root {
	defer fetch.CatchPanic("FindNextSibling()")
	nextSibling := r.Pointer.NextSibling
	if nextSibling == nil {
		panic("No next sibling found")
	}
	return Root{nextSibling, nextSibling.Data}
}

func (r Root) FindPrevSibling() Root {
	defer fetch.CatchPanic("FindPrevSibling()")
	prevSibling := r.Pointer.PrevSibling
	if prevSibling == nil {
		panic("No previous sibling found")
	}
	return Root{prevSibling, prevSibling.Data}
}

// Finds the next element sibling of the pointer in the DOM
// returning a struct with a pointer to it
func (r Root) FindNextElementSibling() Root {
	defer fetch.CatchPanic("FindNextElementSibling()")
	nextSibling := r.Pointer.NextSibling
	if nextSibling == nil {
		panic("No next element sibling found")
	}
	if nextSibling.Type == html.ElementNode {
		return Root{nextSibling, nextSibling.Data}
	}
	p := Root{nextSibling, nextSibling.Data}
	return p.FindNextElementSibling()
}

// Finds the previous element sibling of the pointer in the DOM
// returning a struct with a pointer to it
func (r Root) FindPrevElementSibling() Root {
	defer fetch.CatchPanic("FindPrevElementSibling()")
	prevSibling := r.Pointer.PrevSibling
	if prevSibling == nil {
		panic("No previous element sibling found")
	}
	if prevSibling.Type == html.ElementNode {
		return Root{prevSibling, prevSibling.Data}
	}
	p := Root{prevSibling, prevSibling.Data}
	return p.FindPrevElementSibling()
}

// Returns an array containing key and values of all attributes
func (r Root) Attrs() map[string]string {
	defer fetch.CatchPanic("Attrs()")
	if r.Pointer.Type != html.ElementNode {
		panic("Not an ElementNode")
	}
	if len(r.Pointer.Attr) == 0 {
		return nil
	}
	return fetch.GetKeyValue(r.Pointer.Attr)
}

// Returns the string inside a non-nested element
func (r Root) Text() string {
	defer fetch.CatchPanic("Text()")
	k := r.Pointer.FirstChild
	if k.Type != html.TextNode {
		panic("First child not a text node")
	}
	if k != nil {
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
