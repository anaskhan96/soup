package fetch

import "golang.org/x/net/html"

func FindOnce(n *html.Node, tag string) (*html.Node, bool) {
	if n.Type == html.ElementNode && n.Data == tag {
		return n, true
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		p, q := FindOnce(c, tag)
		if q != false {
			return p, q
		}
	}
	return nil, false
}
