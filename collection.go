package temple

import "errors"

type List []interface{}

func NewList(vals ...interface{}) List {
	return List(vals)
}

type Set map[interface{}]bool

func NewSet(vals ...interface{}) Set {
	s := make(Set)
	for _, v := range vals {
		s[v] = true
	}
	return s
}

func Contains(v interface{}, c interface{}) (bool, error) {
	var in bool
	switch a := c.(type) {
	case List:
		for _, e := range a {
			if e == v {
				in = true
				break
			}
		}
	case Set:
		_, ok := a[v]
		in = ok
	default:
		return false, errors.New("invalid collection")
	}
	return in, nil
}
