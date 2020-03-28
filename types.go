package temple

import "github.com/spf13/cast"

func ToInt(v interface{}) (int, error) {
	return cast.ToIntE(v)
}

func ToInt8(v interface{}) (int8, error) {
	return cast.ToInt8E(v)
}

func ToInt16(v interface{}) (int16, error) {
	return cast.ToInt16E(v)
}

func ToInt32(v interface{}) (int32, error) {
	return cast.ToInt32E(v)
}

func ToInt64(v interface{}) (int64, error) {
	return cast.ToInt64E(v)
}

func ToUint(v interface{}) (uint, error) {
	return cast.ToUintE(v)
}

func Touint8(v interface{}) (uint8, error) {
	return cast.ToUint8E(v)
}

func ToUint16(v interface{}) (uint16, error) {
	return cast.ToUint16E(v)
}

func ToUint32(v interface{}) (uint32, error) {
	return cast.ToUint32E(v)
}

func ToUint64(v interface{}) (uint64, error) {
	return cast.ToUint64E(v)
}

func ToFloat32(v interface{}) (float32, error) {
	return cast.ToFloat32E(v)
}

func ToFloat64(v interface{}) (float64, error) {
	return cast.ToFloat64E(v)
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

func ToInt8Slice(vals []interface{}) ([]int8, error) {
	out := make([]int8, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = ToInt8(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func ToInt16Slice(vals []interface{}) ([]int16, error) {
	out := make([]int16, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = ToInt16(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func ToInt32Slice(vals []interface{}) ([]int32, error) {
	out := make([]int32, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = ToInt32(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func ToInt64Slice(vals []interface{}) ([]int64, error) {
	out := make([]int64, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = ToInt64(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func ToUintSlice(vals []interface{}) ([]uint, error) {
	out := make([]uint, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = ToUint(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func Touint8Slice(vals []interface{}) ([]uint8, error) {
	out := make([]uint8, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = Touint8(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func ToUint16Slice(vals []interface{}) ([]uint16, error) {
	out := make([]uint16, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = ToUint16(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func ToUint32Slice(vals []interface{}) ([]uint32, error) {
	out := make([]uint32, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = ToUint32(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func ToUint64Slice(vals []interface{}) ([]uint64, error) {
	out := make([]uint64, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = ToUint64(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func ToFloat32Slice(vals []interface{}) ([]float32, error) {
	out := make([]float32, len(vals))
	var err error
	for i, v := range vals {
		out[i], err = ToFloat32(v)
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

func ToStringSlice(vals []interface{}) []string {
	out := make([]string, len(vals))
	for i, v := range vals {
		out[i] = ToString(v)
	}
	return out
}
