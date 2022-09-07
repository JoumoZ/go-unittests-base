package Three_Sum

import (
	"reflect"
	"testing"
)

//test table driven
var tests = []struct {
	name   string
	nums   []int
	result [][]int
}{
	{"1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 0, 1}, {-1, 2, -1}}},
	{"2", []int{0, 1, 1}, [][]int{}},
	{"3", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
	{"4", []int{1, -1, -1, 0}, [][]int{{-1, 0, 1}}},
}

func TestCase(t *testing.T) {
	t.Parallel()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := ThreeSum(tc.nums)
			if !reflect.DeepEqual(actual, tc.result) {
				t.Errorf("result not equal: actual: %+v\nexpected:%+v", actual, tc.result)
			}
		})
	}
}
