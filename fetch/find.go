package fetch

import "golang.org/x/net/html"

func FindOnce(n *html.Node, args []string, uni bool) (*html.Node, bool, bool) {
	if uni == true {
		if n.Type == html.ElementNode && n.Data == args[0] {
			if len(args)>1 {
				for i := 0; i < len(n.Attr); i++ {
					if n.Attr[i].Key == args[1] && n.Attr[i].Val == args[2] {
						return n, true, true
					}
				}
			} else {
				return n,true,true
			}
		}
	}
	uni = true
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		p, q, _ := FindOnce(c, args, true)
		if q != false {
			return p, q, true
		}
	}
	return nil, false, true
}
