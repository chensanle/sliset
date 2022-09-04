package sliset

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	type args struct {
		base     []int
		compared []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "compared empty", args: args{base: []int{1, 2, 3}, compared: []int{}}, want: []int{1, 2, 3}},
		{name: "base empty", args: args{base: []int{}, compared: []int{1, 2, 3}}, want: []int{}},
		{name: "default", args: args{base: []int{1, 2, 3}, compared: []int{1, 2}}, want: []int{3}},
		{name: "difference", args: args{base: []int{1, 2, 3}, compared: []int{1, 2, 4}}, want: []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Difference(tt.args.base, tt.args.compared); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	type args struct {
		base     []int
		compared []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{base: []int{1, 2, 3}, compared: []int{}}, want: []int{}},
		{name: "", args: args{base: []int{}, compared: []int{1, 2, 3}}, want: []int{}},
		{name: "", args: args{base: []int{1, 2, 3}, compared: []int{1, 2}}, want: []int{1, 2}},
		{name: "", args: args{base: []int{1, 2, 3}, compared: []int{1, 2, 4}}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.base, tt.args.compared); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniq(t *testing.T) {
	type args struct {
		base []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{base: []int{}}, want: []int{}},
		{name: "", args: args{base: []int{1}}, want: []int{1}},
		{name: "", args: args{base: []int{1, 1, 1, 1, 1}}, want: []int{1}},
		{name: "", args: args{base: []int{1, 2, 3, 1, 2, 3}}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uniq(tt.args.base); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	type args struct {
		base     []int
		compared []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{base: []int{1, 2, 3}, compared: []int{}}, want: []int{1, 2, 3}},
		{name: "", args: args{base: []int{}, compared: []int{1, 2, 3}}, want: []int{1, 2, 3}},
		{name: "", args: args{base: []int{1, 2, 3}, compared: []int{1, 2}}, want: []int{1, 2, 3}},
		{name: "", args: args{base: []int{1, 2, 3}, compared: []int{1, 2, 4}}, want: []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.base, tt.args.compared); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		base []int
		elem int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{base: []int{}, elem: 1}, want: false},
		{name: "", args: args{base: []int{1}, elem: 1}, want: true},
		{name: "", args: args{base: []int{1, 1, 1, 1, 1}, elem: 1}, want: true},
		{name: "", args: args{base: []int{1, 2, 3, 1, 2, 3}, elem: 1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.base, tt.args.elem); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscard(t *testing.T) {
	type args struct {
		base []int
		elem int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{base: []int{}, elem: 1}, want: []int{}},
		{name: "", args: args{base: []int{1}, elem: 1}, want: []int{}},
		{name: "", args: args{base: []int{1, 1, 1, 1, 1}, elem: 1}, want: []int{}},
		{name: "", args: args{base: []int{1, 2, 3, 1, 2, 3}, elem: 1}, want: []int{2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Discard(tt.args.base, tt.args.elem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Discard() = %v, want %v", got, tt.want)
			}
		})
	}
}
