## Variables

### Declaration
```go
var i int = 5
// or
var i = 5

i := 5
// or
var i int = 5


```

### Multiple Variable Assignment

```go
a, b := 1, 2
```

### Public vs Private Variables
```go
package demo

// These are all private variables, and cannot be accessed outside of package demo
var i int
count := 5

// These are public variables
var Name string
Age := 24
```
### Variable Declaration Blocks
```go
var (
  // we can declare variables together in a single `var` block
  name = "John Doe"
  age = 30
  // we can assign multiple values at once
  isLorem, isIpsum = false, true
  // we can declare variables without assigning values
  info string
)
```

