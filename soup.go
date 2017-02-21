package soup

import (
	//"io"
	"golang.org/x/net/html"
	"io/ioutil"
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

func HTMLParse(s string) (*html.Node, error) {
	r, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return nil, err
	}
	return r, nil
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
