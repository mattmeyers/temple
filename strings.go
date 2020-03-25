package temple

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ToString converts a provided value to a string. The
// %v verb is used with fmt.Sprintf to perform the
// conversion for all types except float32 and float64.
// In this case %f is used. An optional precision can be
// passed in this case. Note that the precision must be
// passed as the first argument. For more control over
// the formatting, use the printf builtin.
func ToString(v interface{}, args ...interface{}) string {
	// When an argument is provided, the piped value is passed
	// as the last parameter. In this case, we want to swap
	// these values so that we can do
	//      3.450 | ToString 2
	if len(args) > 0 {
		v, args[0] = args[0], v
	}

	var s string
	switch v.(type) {
	case float64, float32:
		fStr := "%f"
		if len(args) > 0 {
			fStr = fmt.Sprintf("%%.%df", args[0].(int))
		}

		s = fmt.Sprintf(fStr, v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	strconv.Itoa(1)
	return strings.ToLower(s)
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
