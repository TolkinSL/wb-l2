package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFields(s string) map[int]struct{} {
	result := make(map[int]struct{})

	parts := strings.Split(s, ",")

	for _, p := range parts {
		if strings.Contains(p, "-") {
			bounds := strings.Split(p, "-")
			if len(bounds) != 2 {
				continue
			}

			start, err1 := strconv.Atoi(bounds[0])
			end, err2 := strconv.Atoi(bounds[1])

			if err1 != nil || err2 != nil || start > end {
				continue
			}

			for i := start; i <= end; i++ {
				result[i] = struct{}{}
			}
		} else {
			num, err := strconv.Atoi(p)
			if err != nil {
				continue
			}
			result[num] = struct{}{}
		}
	}

	return result
}

func main() {
	fieldsFlag := flag.String("f", "", "flag to select")
	delimiter := flag.String("d", "\t", "delimiter")
	separated := flag.Bool("s", false, "only lines with delimiter")

	flag.Parse()

	if *fieldsFlag == "" {
		fmt.Fprintln(os.Stderr, "flag -f required")
		os.Exit(1)
	}

	fields := parseFields(*fieldsFlag)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		parts := strings.Split(line, *delimiter)

		var out []string

		for i, part := range parts {
			if _, ok := fields[i+1]; ok {
				out = append(out, part)
			}
		}

		if len(out) > 0 {
			fmt.Println(strings.Join(out, *delimiter))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}