package main

import (
	"net/url"
	"strings"
)

func (w *Worker) resolveURL(base, ref string) string {
	u, err := url.Parse(ref)
	if err != nil {
		return ""
	}

	b, err := url.Parse(base)
	if err != nil {
		return ""
	}

	return b.ResolveReference(u).String()
}

func (w *Worker) makeFilePath(link string) string {
	u, _ := url.Parse(link)

	p := u.Path
	if p == "" || strings.HasSuffix(p, "/") {
		p += "index.html"
	}

	return "output/" + u.Host + p
}