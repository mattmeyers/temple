package temple

import (
	"errors"
	"math"
)

func Max(arg1 interface{}, arg2 ...interface{}) (interface{}, error) {
	switch v := arg1.(type) {
	case int, []int:
		return IntMax(arg1, arg2...)
	case uint, []uint:
		return UintMax(arg1, arg2...)
	case float64, []float64:
		return FloatMax(arg1, arg2...)
	case []interface{}:
		if len(v) == 0 {
			return nil, errors.New("empty slice provided")
		}
		switch v[0].(type) {
		case int:
			return IntMax(arg1, arg2...)
		case uint:
			return UintMax(arg1, arg2...)
		case float64:
			return FloatMax(arg1, arg2...)
		}
	case Slice:
		if len(v) == 0 {
			return nil, errors.New("empty slice provided")
		}
		switch v[0].(type) {
		case int:
			return IntMax(arg1, arg2...)
		case uint:
			return UintMax(arg1, arg2...)
		case float64:
			return FloatMax(arg1, arg2...)
		}
	}
	return nil, errors.New("invalid type")
}

func IntMax(arg1 interface{}, arg2 ...interface{}) (int, error) {
	vals, err := parseIntArgs(arg1, arg2...)
	if err != nil {
		return 0, err
	} else if len(vals) == 0 {
		return 0, errors.New("empty slice")
	}

	max := vals[0]
	for _, i := range vals {
		if i > max {
			max = i
		}
	}
	return max, nil
}

func IntMin(arg1 interface{}, arg2 ...interface{}) (int, error) {
	vals, err := parseIntArgs(arg1, arg2...)
	if err != nil {
		return 0, err
	} else if len(vals) == 0 {
		return 0, errors.New("empty slice")
	}

	min := vals[0]
	for _, i := range vals {
		if i < min {
			min = i
		}
	}
	return min, nil
}

func parseIntArgs(arg1 interface{}, arg2 ...interface{}) ([]int, error) {
	var vals []int
	var err error

	switch v := arg1.(type) {
	case int:
		vals, err = ToIntSlice(append([]interface{}{arg1}, arg2...))
	case Slice:
		vals, err = ToIntSlice(v)
	case []interface{}:
		vals, err = ToIntSlice(v)
	case []int:
		vals = v
	default:
		err = errors.New("invalid type")
	}

	return vals, err
}

func UintMax(arg1 interface{}, arg2 ...interface{}) (uint, error) {
	vals, err := parseUintArgs(arg1, arg2...)
	if err != nil {
		return 0, err
	} else if len(vals) == 0 {
		return 0, errors.New("empty slice")
	}

	max := vals[0]
	for _, i := range vals {
		if i > max {
			max = i
		}
	}
	return max, nil
}

func UintMin(arg1 interface{}, arg2 ...interface{}) (uint, error) {
	vals, err := parseUintArgs(arg1, arg2...)
	if err != nil {
		return 0, err
	} else if len(vals) == 0 {
		return 0, errors.New("empty slice")
	}

	min := vals[0]
	for _, i := range vals {
		if i < min {
			min = i
		}
	}
	return min, nil
}

func parseUintArgs(arg1 interface{}, arg2 ...interface{}) ([]uint, error) {
	var vals []uint
	var err error

	switch v := arg1.(type) {
	case uint:
		vals, err = ToUintSlice(append([]interface{}{arg1}, arg2...))
	case Slice:
		vals, err = ToUintSlice(v)
	case []interface{}:
		vals, err = ToUintSlice(v)
	case []uint:
		vals = v
	default:
		err = errors.New("invalid type")
	}

	return vals, err
}

func FloatMax(arg1 interface{}, arg2 ...interface{}) (float64, error) {
	vals, err := parseFloatArgs(arg1, arg2...)
	if err != nil {
		return 0, err
	} else if len(vals) == 0 {
		return 0, errors.New("empty slice")
	}

	max := vals[0]
	for _, i := range vals {
		max = math.Max(max, i)
	}
	return max, nil
}

func FloatMin(arg1 interface{}, arg2 ...interface{}) (float64, error) {
	vals, err := parseFloatArgs(arg1, arg2...)
	if err != nil {
		return 0, err
	} else if len(vals) == 0 {
		return 0, errors.New("empty slice")
	}

	min := vals[0]
	for _, i := range vals {
		min = math.Min(min, i)
	}
	return min, nil
}

func parseFloatArgs(arg1 interface{}, arg2 ...interface{}) ([]float64, error) {
	var vals []float64
	var err error

	switch v := arg1.(type) {
	case float64, float32, int:
		vals, err = ToFloat64Slice(append([]interface{}{arg1}, arg2...))
	case Slice:
		vals, err = ToFloat64Slice(v)
	case []interface{}:
		vals, err = ToFloat64Slice(v)
	case []float64:
		vals = v
	default:
		err = errors.New("invalid type")
	}

	return vals, err
}

func Ceil(f float64) float64 {
	return math.Ceil(f)
}

func Floor(f float64) float64 {
	return math.Floor(f)
}

func Mod(x float64, y float64) float64 {
	return math.Mod(x, y)
}

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Sum(x float64, vals ...float64) float64 {
	for _, v := range vals {
		x += v
	}
	return x
}

func Diff(x float64, vals ...float64) float64 {
	for _, v := range vals {
		x -= v
	}
	return x
}

func Mul(x float64, vals ...float64) float64 {
	for _, v := range vals {
		x *= v
	}
	return x
}

func Div(x float64, vals ...float64) float64 {
	for _, v := range vals {
		if v == 0 {
			panic("temple: division by zero")
		}

		x /= v
	}
	return x
}
