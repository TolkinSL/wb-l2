package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type Worker struct {
	visited map[string]bool
	domain  string
}

func NewWorker(domain string) *Worker {
	return &Worker{
		visited: make(map[string]bool),
		domain:  domain,
	}
}

func (w *Worker) Download(link string, depth, maxDepth int) error {
	if depth > maxDepth {
		return nil
	}

	if w.visited[link] {
		return nil
	}
	w.visited[link] = true

	fmt.Println("download:", link)

	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	filePath := w.makeFilePath(link)
	err = os.MkdirAll(path.Dir(filePath), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return nil
	}

	return w.parseAndDownload(link, data, depth, maxDepth)
}