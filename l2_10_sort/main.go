package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/TolkinSL/wb-l2/l2_10_sort/sortlines"
	"os"
)

func main() {
	cfg := parseFlags()

	lines, err := readLines()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	lines = sortlines.SortLines(lines, cfg)

	if cfg.Unique {
		lines = sortlines.Unique(lines)
	}

	for _, l := range lines {
		fmt.Println(l)
	}
}

func parseFlags() sortlines.Config {
	var cfg sortlines.Config

	flag.IntVar(&cfg.Column, "k", 0, "column number")
	flag.BoolVar(&cfg.Numeric, "n", false, "numeric sort")
	flag.BoolVar(&cfg.Reverse, "r", false, "reverse sort")
	flag.BoolVar(&cfg.Unique, "u", false, "unique lines")

	flag.Parse()

	return cfg
}

func readLines() ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
