package temple

import (
	"errors"
	"strings"
)

// ToString converts a provided value to a string. The
// %v verb is used with fmt.Sprintf to perform the
// conversion for all types except float32 and float64.
// In this case %f is used. An optional precision can be
// passed in this case. Note that the precision must be
// passed as the first argument. For more control over
// the formatting, use the printf builtin.
// func ToString(v interface{}, args ...interface{}) string {
// 	// When an argument is provided, the piped value is passed
// 	// as the last parameter. In this case, we want to swap
// 	// these values so that we can do
// 	//      3.450 | ToString 2
// 	if len(args) > 0 {
// 		v, args[0] = args[0], v
// 	}

// 	var s string
// 	switch v.(type) {
// 	case float64, float32:
// 		fStr := "%f"
// 		if len(args) > 0 {
// 			fStr = fmt.Sprintf("%%.%df", args[0].(int))
// 		}

// 		s = fmt.Sprintf(fStr, v)
// 	default:
// 		s = fmt.Sprintf("%v", v)
// 	}
// 	return s
// }

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
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
	l := len(s)

	if l < 4 {
		return 0
	}

	if s[0] == '-' {
		l--
	}

	return (l - 1) / 3
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

func Join(sep string, a interface{}) (out string, err error) {
	switch s := a.(type) {
	case []string:
		out = strings.Join(s, sep)
	case []interface{}:
		sl, err := ToStringSlice(s)
		if err != nil {
			return "", err
		}
		out = strings.Join(sl, sep)
	case List:
		sl, err := ToStringSlice(s)
		if err != nil {
			return "", err
		}
		out = strings.Join(sl, sep)
	default:
		return "", errors.New("slice of strings required")
	}
	return out, nil
}

// FormatMask applies a mask to the provided string. A mask is a string containing
// the '#' rune. When processed, each '#' will be replaced by the corresponding
// rune in the provided string. For example, to format a phone number, use
//		FormatMask(`(###) ###-####`, "5551234567") -> "(555) 123-4567"
//
// To include a literal '#' in the mask, use a raw string and precede the '#'
// with a '\'. For example:
//		{{ FormatMask `\##-##` "123" }} -> "#1-23"
//
// This function expects the number of '#' runes in the mask to match the number of
// runes in the provided string. An error will be returned when these lengths
// do not match.
func FormatMask(mask string, str string) (string, error) {
	if len(str) == 0 {
		return mask, nil
	}

	mRunes, sRunes := []rune(mask), []rune(str)
	if len(mRunes) < len(sRunes) {
		return "", errors.New("mask too short for string")
	}

	j := 0
	var escape bool
	for i := 0; i < len(mRunes); i++ {
		switch mRunes[i] {
		case '#':
			if escape {
				mRunes = append(mRunes[0:i-1], mRunes[i:]...)
				escape = false
				i--
				continue
			}
		case '\\':
			escape = true
			continue
		default:
			escape = false
			continue
		}

		if j >= len(sRunes) {
			return "", errors.New("too few string characters for mask")
		}

		mRunes[i] = sRunes[j]
		j++
	}

	if j != len(sRunes) {
		return "", errors.New("unused string characters")
	}

	return string(mRunes), nil
}
