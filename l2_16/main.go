package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run main.go <url>")
	}

	startURL := os.Args[1]

	u, err := url.Parse(startURL)
	if err != nil {
		log.Fatal(err)
	}

	w := NewWorker(u.Host)

	err = w.Download(startURL, 0, 2) // depth = 2
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}