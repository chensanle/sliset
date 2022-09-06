sliset: a series basic set operations for slice type
===
[![Doc of Sliset][4]][3]
[![Go Report Card][2]][1]
[![Codecov][6]][5]
[![golang-lint][8]][7]

基于泛型的 slice 伪集合操作，封装常见的 union、interaction 和 difference 等函数。

## Installation

Standard  `go get`:

```
$  go get -v -u github.com/chensanle/sliset
```

## Usage & Example
```go
// Difference res = base - compared
func Difference(base, compared []E1) []E1

// Intersection res = base ∩ compared
func Intersection(base, compared []E1) []E1

// Union res = base U compared
func Union(base, compared []E1) []E1

// Uniq remove duplicate elements from the base.
func Uniq(base []E1) []E1
```

###  example 1:  Difference()
```go
r1 := sliset.Difference([]string{"apple", "fb", "ali"}, []string{"apple"})
fmt.Println("output: ", r1)
// output: []string{"ali"}

r2 := sliset.Difference([]int{1949, 1997, 2008}, []int{2008})
fmt.Println("output: ", r2)
// output: []string{1949, 1997}
```

###  example 2:  Uniq()
```go
r1 := sliset.Uniq([]string{"apple", "fb", "ali", "apple"})
fmt.Println("output: ", r1)
// output: []string{"apple", "fb", "ali"} 

r2 := sliset.Uniq([]int{1949, 1997, 2008, 1949})
fmt.Println("output: ", r2)
// output: []string{1949, 1997, 2008}
```

[1]: <https://goreportcard.com/report/github.com/chensanle/sliset> "Go Report Card Link"
[2]: <https://goreportcard.com/badge/github.com/chensanle/sliset> "Go Report Card Badge"
[3]: <https://pkg.go.dev/github.com/chensanle/sliset?tab=doc> "Doc of Sliset Link"
[4]: <https://img.shields.io/badge/go.dev-doc-007d9c?style=flat-square&logo=read-the-docs> "Doc of Sliset Badge"

[5]: <https://codecov.io/gh/chensanle/sliset> "Codecov Link"
[6]: <https://codecov.io/gh/chensanle/sliset/branch/main/graph/badge.svg?token=BAR2JNMTUJ> "Codecov Badge"

[7]: <https://github.com/chensanle/sliset/actions/workflows/golangci-lint.yml> "Link"
[8]: <https://github.com/chensanle/sliset/actions/workflows/golangci-lint.yml/badge.svg> "Badge"