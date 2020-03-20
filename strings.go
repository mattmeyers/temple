package temple

import (
	"errors"
	"fmt"
	"strings"
)

// ToString converts a provided value to a string. The
// %v verb is used with fmt.Sprintf to perform the
// conversion.
func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

// Commas adds a comma after every three digits.
func Commas(s string) (string, error) {
	if !IsNumeric(s) {
		return "", errors.New("non numeric string")
	}

	if s[0] == '.' {
		return s, nil
	}

	parts := strings.Split(s, ".")

	n := numCommas(parts[0])
	if n == 0 {
		return s, nil
	}

	b := []byte(parts[0])
	t := make([]byte, len(b)+n)
	copy(t, b)
	b = t

	k := 0
	for j := len(b) - 1; j > 0; j-- {
		if k == 3 {
			b[j] = ','
			n--
			k = 0
		} else {
			b[j] = b[j-n]
			k++
		}
	}
	parts[0] = string(b)

	return strings.Join(parts, "."), nil
}

// IsNumeric determines if the provided string is a
// valid number. Valid numbers can contains commas to
// the left of the decimal point.
func IsNumeric(s string) bool {
	b := []byte(s)
	if len(b) == 0 {
		return false
	}

	if b[0] == '-' && len(b) > 1 {
		b = b[1:]
	}

	dPt := false
	for _, d := range b {
		if '0' <= d && d <= '9' {
			continue
		} else if d == ',' && !dPt {
			continue
		} else if d == '.' && !dPt {
			dPt = true
		} else {
			return false
		}
	}

	return true
}

func numCommas(s string) int {
	n := 0
	switch l := len(s); {
	case l < 4:
		break
	case s[0] == '-':
		l--
		fallthrough
	default:
		n = (l - 1) / 3
	}
	return n
}
