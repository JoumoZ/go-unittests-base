package Two_Sum

import (
	"reflect"
	"testing"
)

//test table driven
var tests = []struct {
	name   string
	nums   []int
	target int
	result []int
}{
	{"1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
	{"2", []int{3, 2, 4}, 6, []int{1, 2}},
	{"3", []int{3, 3}, 7, []int{0, 1}}, //expected target = 6
}

func TestTwoSum(t *testing.T) {
	t.Parallel()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := TwoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(actual, tc.result) {
				t.Error("result not equal")
			}
		})
	}
}
