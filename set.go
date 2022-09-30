package sliset

import (
	"golang.org/x/exp/slices"
)

// Contains returns the result of elem whether exists in base,
// if existed, return true, or false.
func Contains[E1 Int | Float | string](base []E1, elem E1) bool {
	if Index(base, elem) >= 0 {
		return true
	}
	return false
}

// Index returns the index of the first occurrence of elem in base,
// or -1 if not present.
func Index[E1 Int | Float | string](base []E1, elem E1) int {
	return slices.Index(base, elem)
}

// Difference res = base - compared
func Difference[E1 Int | Float | string](base, compared []E1) []E1 {
	if len(base) <= 0 {
		return make([]E1, 0)
	}
	comparedSet := S2Set(compared)

	newRes := make([]E1, 0)
	for _, val := range base {
		if comparedSet[val] {
			continue
		}
		newRes = append(newRes, val)
	}

	return Uniq(newRes)
}

// Intersection res = base ∩ compared
func Intersection[E1 Int | Float | string](base, compared []E1) []E1 {
	if len(base) <= 0 || len(compared) <= 0 {
		return make([]E1, 0)
	}

	rightSet := S2Set(compared)
	newRes := make([]E1, 0)
	for _, val := range base {
		if rightSet[val] {
			newRes = append(newRes, val)
		}
	}

	return Uniq(newRes)
}

// Union res = base U compared
func Union[E1 Int | Float | string](base, compared []E1) []E1 {
	res := make([]E1, len(base)+len(compared))
	i := 0
	for _, val := range base {
		res[i] = val
		i++
	}
	for _, val := range compared {
		res[i] = val
		i++
	}

	return Uniq(res)
}

// Uniq remove duplicate elements from the base.
func Uniq[E1 Int | Float | string](base []E1) []E1 {
	if len(base) <= 0 {
		return make([]E1, 0)
	}

	res := make([]E1, 0)
	cache := make(map[E1]struct{})

	for _, val := range base {
		if _, ok := cache[val]; ok {
			continue
		}
		res = append(res, val)
		cache[val] = struct{}{}
	}

	return res
}

// Discard remove specified elem from the base.
func Discard[E1 Int | Float | string](base []E1, elem E1) []E1 {
	res := make([]E1, 0)
	for _, val := range base {
		if val == elem {
			continue
		}
		res = append(res, val)
	}
	return Uniq(res)
}

// S2Set turn slice to set. no used map[E1]struct{} to save memory because
// the bool may reduce the program's complexity.
func S2Set[E1 Int | Float | string](base []E1) map[E1]bool {
	m := make(map[E1]bool)
	for _, val := range base {
		m[val] = true
	}
	return m
}

// S2Map turn slice to map, the key is custom by getKey()
func S2Map[E1 any, E2 comparable](base []E1, getKey func(E1) E2) map[E2]E1 {
	m := make(map[E2]E1)
	for i, val := range base {
		m[getKey(val)] = base[i]
	}

	return m
}

func Copy[E1 Int | Float | string](base []E1) []E1 {
	res := make([]E1, len(base))

	for i, _ := range base {
		res[i] = base[i]
	}
	return res
}

// Discard

// IsDisJoint 判断两个集合是否包含相同元素
func IsDisJoint[E1 Int | Float | string](base, compared []E1) bool {
	if len(base) <= 0 || len(compared) <= 0 {
		return false
	}
	comparedSet := S2Set(compared)
	for _, val := range base {
		if comparedSet[val] {
			return true
		}
	}
	return false
}

// IsSubset Is the compared slice a sub set of base
//func IsSubset[E1 Int | Float | string](base, compared []E1) bool {
//	return false
//}
//
//func Remove[E1 Int | Float | string](base []E1, elem E1) int {
//	return 0
//}
