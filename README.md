# LINQ for Go

### usage
```golang

import "github.com/moamlrh/linqit"

func main() {
    numbers := []int{1, 3, 3, 4, 5, 6}
    
    result := linqit.Array(numbers).
        Where(func(i int) bool { return i > 2 }).
        Distinct(func(a, b int) bool { return a == b }).
        OrderBy(func(a, b int) bool { return a > b }).
        ToSlice()
    
    fmt.Println(result) // output: [6 5 4 3]
}
```
