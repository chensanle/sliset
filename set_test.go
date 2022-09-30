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

func TestS2Map(t *testing.T) {
	type user struct {
		Uid      int
		Nickname string
	}

	userList := make([]*user, 0)
	userList = []*user{{Uid: 1, Nickname: "jack"}, {Uid: 2, Nickname: "perri"}}

	m := S2Map(userList, func(user *user) int { return user.Uid })
	for key, val := range m {
		t.Log(key, ":", val)
	}
}

func TestS2Map1(t *testing.T) {
	type user struct {
		Uid      int
		Nickname string
	}
	type unit[E1 any, E2 comparable] struct {
		name   string
		base   []E1
		getKey func(E1) E2
		want   map[E2]E1
	}
	tests := []unit[user, int]{
		{name: "int-first", base: []user{{Uid: 1, Nickname: "jack"}, {Uid: 2, Nickname: "perri"}}, getKey: func(u user) int { return u.Uid }, want: map[int]user{1: {Uid: 1, Nickname: "jack"}, 2: {Uid: 2, Nickname: "perri"}}},
		{name: "int-seconds", base: []user{{Uid: 1, Nickname: "jack"}}, getKey: func(u user) int { return u.Uid }, want: map[int]user{1: {Uid: 1, Nickname: "jack"}}},
	}
	tests2 := []unit[user, string]{
		{name: "string-first", base: []user{{Uid: 1, Nickname: "jack"}, {Uid: 2, Nickname: "perri"}}, getKey: func(u user) string { return u.Nickname }, want: map[string]user{"jack": {Uid: 1, Nickname: "jack"}, "perri": {Uid: 2, Nickname: "perri"}}},
		{name: "string-seconds", base: []user{{Uid: 1, Nickname: "jack"}}, getKey: func(u user) string { return u.Nickname }, want: map[string]user{"jack": {Uid: 1, Nickname: "jack"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := S2Map(tt.base, tt.getKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("S2Map() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := S2Map(tt.base, tt.getKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("S2Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDisJoint(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			base     []int
			compared []int
		}
		want bool
	}{
		{
			name: "first",
			args: struct {
				base     []int
				compared []int
			}{base: []int{1, 2, 3, 4}, compared: []int{1, 2, 3, 4}},
			want: true,
		},
		{
			name: "seconds",
			args: struct {
				base     []int
				compared []int
			}{base: []int{1, 2, 3, 4}, compared: []int{5, 6, 7, 8}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDisJoint(tt.args.base, tt.args.compared); got != tt.want {
				t.Errorf("IsDisJoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
