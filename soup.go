package soup

import (
	"io"
	"io/ioutil"
	"net/http"
)

func Make(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

/*func Prettify(r io.ReadCloser) (string, error) {
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
}*/
