package main

import (
	"bytes"
	"net/url"

	"golang.org/x/net/html"
)

func (w *Worker) parseAndDownload(base string, body []byte, depth, maxDepth int) error {
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "href" || attr.Key == "src" {
					link := w.resolveURL(base, attr.Val)

					if link == "" {
						continue
					}

					u, err := url.Parse(link)
					if err != nil {
						continue
					}

					if u.Host != w.domain {
						continue
					}

					w.Download(link, depth+1, maxDepth)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
	return nil
}