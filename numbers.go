package temple

func SliceMax(vals []float64) float64 {
	if len(vals) == 0 {
		return 0
	}

	m := vals[0]
	for _, v := range vals {
		if v > m {
			m = v
		}
	}

	return m
}

func IntSliceMax(vals []int) int {
	if len(vals) == 0 {
		return 0
	}

	m := vals[0]
	for _, v := range vals {
		if v > m {
			m = v
		}
	}

	return m
}

func UintSliceMax(vals []uint) uint {
	if len(vals) == 0 {
		return 0
	}

	m := vals[0]
	for _, v := range vals {
		if v > m {
			m = v
		}
	}

	return m
}

func Max(v float64, args ...float64) float64 {
	if len(args) == 0 {
		return v
	}

	return SliceMax(append(args, v))
}

func IntMax(v int, args ...int) int {
	if len(args) == 0 {
		return v
	}

	return IntSliceMax(append(args, v))
}

func UintMax(v uint, args ...uint) uint {
	if len(args) == 0 {
		return v
	}

	return UintSliceMax(append(args, v))
}

func SliceMin(vals []float64) float64 {
	if len(vals) == 0 {
		return 0
	}

	m := vals[0]
	for _, v := range vals {
		if v < m {
			m = v
		}
	}

	return m
}

func IntSliceMin(vals []int) int {
	if len(vals) == 0 {
		return 0
	}

	m := vals[0]
	for _, v := range vals {
		if v < m {
			m = v
		}
	}

	return m
}

func UintSliceMin(vals []uint) uint {
	if len(vals) == 0 {
		return 0
	}

	m := vals[0]
	for _, v := range vals {
		if v < m {
			m = v
		}
	}

	return m
}

func Min(f float64, args ...float64) float64 {
	if len(args) == 0 {
		return f
	}

	return SliceMin(append(args, f))
}

func IntMin(f int, args ...int) int {
	if len(args) == 0 {
		return f
	}

	return IntSliceMin(append(args, f))
}

func UintMin(f uint, args ...uint) uint {
	if len(args) == 0 {
		return f
	}

	return UintSliceMin(append(args, f))
}
