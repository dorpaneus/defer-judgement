package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"classic example", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"reuses earlier index", []int{3, 2, 4}, 6, []int{1, 2}},
		{"same number twice", []int{3, 3}, 6, []int{0, 1}},
		{"no pair returns nil", []int{1, 2, 3}, 100, nil},
		{"empty input", []int{}, 5, nil},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TwoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TwoSum(%v, %d) = %v; want %v", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}
