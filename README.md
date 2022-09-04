sliset: a series basic set operations for slice type
===

基于泛型的 slice 集合操作，封装常见的 union、interaction 和 difference 等。

## Installation

Standard  `go get`:

```
$  go get -v -u github.com/chensanle/slices
```

## Usage & Example


###  example 1:  Difference()
```go
r1 := sliset.Difference([]string{"apple", "fb", "ali"}, []string{"apple"})
fmt.Println("res=", r1)
// output: []string{"ali"}

r2 := sliset.Difference([]string{1949, 1997, 2008}, []string{2008})
fmt.Println("res=", r2)
// output: []string{1949, 1997}
```

###  example 2:  Uniq()
```go
r1 := sliset.Difference([]string{"apple", "fb", "ali", "apple"}, []string{"apple"})
fmt.Println("res=", r1)
// output: []string{"apple", "fb", "ali"} 

r2 := sliset.Difference([]string{1949, 1997, 2008}, []string{2008})
fmt.Println("res=", r2)
// output: []string{1949, 1997, 2008}
```