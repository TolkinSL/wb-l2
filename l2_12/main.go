package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	after      int
	before     int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	pattern    string
}

func main() {
	cfg := parseFlags()

	matcher, err := checkMatch(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	lines, err := readLines()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	processLines(lines, matcher, cfg)
}

func readLines() ([]string, error) {

	var lines []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func parseFlags() Config {
	var cfg Config

	flag.IntVar(&cfg.after, "A", 0, "print N lines after match")
	flag.IntVar(&cfg.before, "B", 0, "print N lines before match")

	context := flag.Int("C", 0, "print N lines of context")

	flag.BoolVar(&cfg.count, "c", false, "count matches")
	flag.BoolVar(&cfg.ignoreCase, "i", false, "ignore case")
	flag.BoolVar(&cfg.invert, "v", false, "invert match")
	flag.BoolVar(&cfg.fixed, "F", false, "fixed string match")
	flag.BoolVar(&cfg.lineNum, "n", false, "print line numbers")

	flag.Parse()

	if *context > 0 {
		cfg.after = *context
		cfg.before = *context
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("pattern required")
		os.Exit(1)
	}

	cfg.pattern = args[0]

	return cfg
}

func checkMatch(cfg Config) (func(string) bool, error) {

	pattern := cfg.pattern

	if cfg.ignoreCase {
		pattern = strings.ToLower(pattern)
	}

	if cfg.fixed {
		return func(line string) bool {

			if cfg.ignoreCase {
				line = strings.ToLower(line)
			}

			return strings.Contains(line, pattern)

		}, nil
	}

	if cfg.ignoreCase {
		pattern = "(?i)" + pattern
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	return func(line string) bool {
		return re.MatchString(line)
	}, nil
}

func processLines(lines []string, matcher func(string) bool, cfg Config) {

	matchCount := 0

	for i := 0; i < len(lines); i++ {

		match := matcher(lines[i])

		if cfg.invert {
			match = !match
		}

		if !match {
			continue
		}

		matchCount++

		if cfg.count {
			continue
		}

		start := i - cfg.before
		if start < 0 {
			start = 0
		}

		end := i + cfg.after
		if end >= len(lines) {
			end = len(lines) - 1
		}

		for j := start; j <= end; j++ {

			if cfg.lineNum {
				fmt.Printf("%d:", j+1)
			}

			fmt.Println(lines[j])
		}
	}

	if cfg.count {
		fmt.Println(matchCount)
	}
}
