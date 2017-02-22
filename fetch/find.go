package fetch

import "golang.org/x/net/html"

func FindOnce(n *html.Node, tag string, uni bool) (*html.Node, bool, bool) {
	if uni == true {
		if n.Type == html.ElementNode && n.Data == tag {
			return n, true, true
		}
	}
	uni = true
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		p, q, _ := FindOnce(c, tag, true)
		if q != false {
			return p, q, true
		}
	}
	return nil, false, true
}
