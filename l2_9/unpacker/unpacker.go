package unpacker

import (
	"fmt"
	"strings"
	"unicode"
)

func Unpack(s string) (string, error) {
	r := []rune(s)

	if len(r) == 0 {
		return "", nil
	}

	if unicode.IsDigit(r[0]) {
		return "", fmt.Errorf("string starts with a digit: %s", string(r))
	}

	var myStr strings.Builder
	var prevRune rune
	var canRepeat bool
	var isEscaped bool

	for _, myRune := range r {
		if myRune == '\\' && !isEscaped {
			isEscaped = true
			continue
		}

		if unicode.IsDigit(myRune) && !isEscaped {
			if !canRepeat {
				return "", fmt.Errorf("digit follows another digit: %s", string(r))
			}

			num := int(myRune - '0')

			for i := 0; i < num-1; i++ {
				myStr.WriteRune(prevRune)
			}

			canRepeat = false
			continue
		}

		myStr.WriteRune(myRune)
		prevRune = myRune
		canRepeat = true
		isEscaped = false
	}

	if isEscaped {
		return "", fmt.Errorf("last escape character \\ is not closed: %s", string(r))
	}

	return myStr.String(), nil
}
