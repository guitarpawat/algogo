package util

import (
	"reflect"
	"testing"
)

func Test_SwapInt(t *testing.T) {
	type args struct {
		slice []int
		i1    int
		i2    int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{"one elem", args{[]int{0}, 0, 0}, []int{0}},
		{"two elem 0,1", args{[]int{0, 1}, 0, 1}, []int{1, 0}},
		{"two elem 1,0", args{[]int{0, 1}, 1, 0}, []int{1, 0}},
		{"three elem 1,0", args{[]int{0, 1, 2}, 1, 0}, []int{1, 0, 2}},
		{"three elem 0,1", args{[]int{0, 1, 2}, 0, 1}, []int{1, 0, 2}},
		{"three elem 1,1", args{[]int{0, 1, 2}, 1, 1}, []int{0, 1, 2}},
		{"three elem 0,2", args{[]int{0, 1, 2}, 0, 2}, []int{2, 1, 0}},
		{"three elem 2,2", args{[]int{0, 1, 2}, 2, 2}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SwapInt(tt.args.slice, tt.args.i1, tt.args.i2)
			if len(tt.expected) != len(tt.args.slice) {
				t.Errorf("length mismatch, expected: %d, but get: %d", len(tt.expected), len(tt.args.slice))
				return
			}
			if !reflect.DeepEqual(tt.expected, tt.args.slice) {
				t.Errorf("slice mismatch\nexpected: %v\nbut get:  %v", tt.expected, tt.args.slice)
				return
			}
		})
	}
}

func Test_ShiftRightAndPutFirstInt(t *testing.T) {
	type args struct {
		slice []int
		index int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{"one elem", args{[]int{0}, 0}, []int{0}},
		{"two elem 0", args{[]int{0, 1}, 0}, []int{0, 1}},
		{"two elem 1", args{[]int{0, 1}, 1}, []int{1, 0}},
		{"three elem 0", args{[]int{0, 1, 2}, 0}, []int{0, 1, 2}},
		{"three elem 1", args{[]int{0, 1, 2}, 1}, []int{1, 0, 2}},
		{"three elem 2", args{[]int{0, 1, 2}, 2}, []int{2, 0, 1}},
		{"four elem 0", args{[]int{0, 1, 2, 3}, 0}, []int{0, 1, 2, 3}},
		{"four elem 1", args{[]int{0, 1, 2, 3}, 1}, []int{1, 0, 2, 3}},
		{"four elem 2", args{[]int{0, 1, 2, 3}, 2}, []int{2, 0, 1, 3}},
		{"four elem 3", args{[]int{0, 1, 2, 3}, 3}, []int{3, 0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShiftRightAndPutFirstInt(tt.args.slice, tt.args.index)
			if len(tt.expected) != len(tt.args.slice) {
				t.Errorf("length mismatch, expected: %d, but get: %d", len(tt.expected), len(tt.args.slice))
				return
			}
			if !reflect.DeepEqual(tt.expected, tt.args.slice) {
				t.Errorf("slice mismatch\nexpected: %v\nbut get:  %v", tt.expected, tt.args.slice)
				return
			}
		})
	}
}
