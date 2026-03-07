package sortlines

import (
	"sort"
	"strconv"
	"strings"
)

type Config struct {
	Column  int
	Numeric bool
	Reverse bool
	Unique  bool
}

func SortLines(lines []string, cfg Config) []string {

	sort.Slice(lines, func(i, j int) bool {

		a := lines[i]
		b := lines[j]

		if cfg.Column > 0 {
			a = getColumn(a, cfg.Column-1)
			b = getColumn(b, cfg.Column-1)
		}

		if cfg.Numeric {
			ai, _ := strconv.ParseFloat(a, 64)
			bi, _ := strconv.ParseFloat(b, 64)

			if cfg.Reverse {
				return ai > bi
			}

			return ai < bi
		}

		if cfg.Reverse {
			return a > b
		}

		return a < b
	})

	return lines
}

func getColumn(line string, k int) string {

	fields := strings.Split(line, "\t")

	if k >= len(fields) {
		return ""
	}

	return fields[k]
}

func Unique(lines []string) []string {

	if len(lines) == 0 {
		return lines
	}

	result := []string{lines[0]}

	for i := 1; i < len(lines); i++ {
		if lines[i] != lines[i-1] {
			result = append(result, lines[i])
		}
	}

	return result
}