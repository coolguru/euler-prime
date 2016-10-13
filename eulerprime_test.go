package eulerprime

import (
	"errors"
	"reflect"
	"testing"
)

type testpair struct {
	values []int
	eprime int64
	err    error
}

var tests = []testpair{
	{[]int{3, 3}, 523, nil},
	{[]int{4, 5}, 24977, nil},
	{[]int{6, 10}, 1063686487, nil},
	{[]int{5, 12}, 487016953141, nil},
	{[]int{3, 69999}, 0, errors.New("Euler prime not found")},
	{[]int{6, 13}, 0, errors.New("Input is outside supported range, try with x <= 5 and y <= 12")},
}

func TestEulerPrime(t *testing.T) {
	for _, pair := range tests {
		v, err := EulerPrime(pair.values[0], pair.values[1])
		if v != pair.eprime || !reflect.DeepEqual(err, pair.err) {
			t.Error(
				"For", pair.values,
				"expected", pair.eprime,
				"got", v, err.Error(),
			)
		}
	}
}
