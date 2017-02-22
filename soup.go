package soup

import (
	"github.com/anaskhan96/soup/fetch"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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

func HTMLParse(s string) Test {
	r, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(s, "<!") {
		return Root{r.FirstChild.NextSibling}
	}
	return Root{r}
}

/* WIP : Find Function */
type Test interface {
	Find(tag string) Test
	Tag() string
	Attrs() []html.Attribute
}

type Root struct {
	Pointer *html.Node
}

func (r Root) Find(tag string) Test {
	temp, ok, _ := fetch.FindOnce(r.Pointer, tag, false)
	if ok == false {
		return nil
	}
	return Root{temp}
}

func (r Root) Tag() string {
	return r.Pointer.Data
}

func (r Root) Attrs() []html.Attribute {
	return r.Pointer.Attr
}

/* ------------------- */

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
