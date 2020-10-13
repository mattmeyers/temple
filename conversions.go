package temple

import "github.com/spf13/cast"

func ToInt(v interface{}) (int, error) {
	return cast.ToIntE(v)
}

func ToFloat64(v interface{}) (float64, error) {
	return cast.ToFloat64E(v)
}

func ToString(v interface{}) (string, error) {
	return cast.ToStringE(v)
}

func ToIntSlice(vals []interface{}) ([]int, error) {
	out := make([]int, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = ToInt(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func ToFloat64Slice(vals []interface{}) ([]float64, error) {
	out := make([]float64, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = ToFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func ToStringSlice(vals []interface{}) ([]string, error) {
	out := make([]string, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = ToString(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
