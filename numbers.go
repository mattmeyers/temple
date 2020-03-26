package temple

import (
	"errors"
	"math"
	"reflect"
)

func Max(arg1 reflect.Value, arg2 ...reflect.Value) (reflect.Value, error) {
	if arg1.Kind() == reflect.Slice || arg1.Kind() == reflect.Array {
		l := arg1.Len()
		if l == 0 {
			return reflect.Value{}, errors.New("empty slice provided")
		}

		out := arg1.Index(0)
		for i := 1; i < l; i++ {
			ok, err := lt(out, arg1.Index(i))
			if err != nil {
				return reflect.Value{}, err
			}

			if ok {
				out = arg1.Index(i)
			}
		}
		return out, nil
	}

	out := arg1
	for _, a := range arg2 {
		ok, err := lt(out, a)
		if err != nil {
			return reflect.Value{}, err
		}
		if ok {
			out = a
		}
	}

	return out, nil
}

func Min(arg1 reflect.Value, arg2 ...reflect.Value) (reflect.Value, error) {
	if arg1.Kind() == reflect.Slice || arg1.Kind() == reflect.Array {
		l := arg1.Len()
		if l == 0 {
			return reflect.Value{}, errors.New("empty slice provided")
		}

		out := arg1.Index(0)
		for i := 1; i < l; i++ {
			ok, err := gt(out, arg1.Index(i))
			if err != nil {
				return reflect.Value{}, err
			}

			if ok {
				out = arg1.Index(i)
			}
		}
		return out, nil
	}

	out := arg1
	for _, a := range arg2 {
		ok, err := gt(out, a)
		if err != nil {
			return reflect.Value{}, err
		}
		if ok {
			out = a
		}
	}

	return out, nil
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
