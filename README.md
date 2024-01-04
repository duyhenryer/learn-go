# learn-go
*Noted:*

Trong Go, Chỉ có public, private 
* <kbd> array </kbd> - Nó là value type
* <kbd> slice </kbd> - Nó là reference type tham chiếu tới mảng. Slice thường được tạo bằng cắt 1 phần của array. Mọi thay đổi trong slice sẽ ảnh hưởng đến mảng gốc và ngược lại.

*Cli*
```go
go mod init
go get -u gopkg.in/yaml.v3

```

```go
	// Chỉ muốn làm việc với value, bỏ qua index và ngược lại. 
	for _, value := range countries {
		fmt.Println(value)
	}
    for index, _ := range countries {
    fmt.Println(index)
    }
```


### Syntax
* <kbd> array </kbd> 
```go
# Declare 1
array_name := [length]Type{item1, item2, item3,...itemN}
a := [5]int{1, 2, 3, 4, 5}

# Declare 2
var a[5]int

# Declare 
```
* <kbd> slice </kbd> 
```go
# Declare 1
[]T
or
[]T{}
or
[]T{value1, value2, value3, ...value n}

var my_slice[]int
var my_slice_1 = []string{"Geeks", "for", "Geeks"}


```
    

## Keywords in Go
### 25 Keywords
```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```
Let’s divide all 25 keywords in 4 groups.
### Declaration
* <kbd> const </kbd>

The `const` keyword is used to introduce a name for a scalar value like 2 or 3.14, etc.
```go
package main

import "fmt"

const pi = 3.14

func main() {
    fmt.Println("Value of pi:", pi)
}
```
* <kbd> var </kbd>

The `var` keyword is used to create the variables in the `Go` language.
```go
package main

import "fmt"

func main() {
	var x int = 5
	fmt.Println("Value of x:", x)
}
```
* <kbd> func </kbd>

The `func` keyword is used to declare a function.
```go
package main

import "fmt"

func greet() {
    fmt.Println("Hello!")
}

func main() {
    greet()
}
```
* <kbd> type </kbd>

We can use the `type` keyword to introduce a new struct type.

```go
package main

import "fmt"

type Celsius float64

func main() {
	var temperature Celsius = 23.5
	fmt.Println("Temperature:", temperature, "Celsius")
}
```
* <kbd> import </kbd>
The `import` keyword is used to import packages.
* <kbd> package </kbd>
The code is grouped as a unit in a package. The `package` keyword is used to define one.
### Composite Types

* <kbd> chan </kbd>
The `chan` keyword is used to define a channel. In `Go`, you are allowed to run parallel pieces of code simultaneously.
* <kbd> interface </kbd>
The `interface` keyword is used to specify a method set. A method set is a list of methods for a type.
* <kbd> map </kbd>

The `map` keyword defines a map type. A map os an unordered collection of the key-value pairs.
```go
package main

import "fmt"

func main() {
    studentGrades := map[string]int{
        "Alice": 90,
        "Bob":   85,
        "Charlie": 95,
    }

    fmt.Println("Bob's Grade:", studentGrades["Bob"])
}
```
* <kbd> struct </kbd>

Struct is a collection of fields. We can use the `struct` keyword followed by the field declarations.
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{"John Doe", 25}
    fmt.Println("Person:", person)
}
```
### Control Flow

* <kbd> break </kbd>
The `break` keyword lets you terminate a loop and continue execution of the rest of the code.
* <kbd> case </kbd> 
This is a form of a `switch` construct. We specify a variable after the switch.
* <kbd> continue </kbd>
The `continue` keyword allows you to go back to the beginning of the ‘for’ loop.
* <kbd> default </kbd>
The `default` statement is optional but you need to use either case or default within the switch statement. 
The switch jumps to the default value if the case does not match the controlling expression.
* <kbd> else </kbd>
The `else` keyword is used with the `if` keyword to execute an alternate statement if a certain condition is false.
* <kbd> fallthrough </kbd>
This keyword is used in a case within a switch statement. When we use this keyword, the next case is entered.
* <kbd> for </kbd>

The `for` keyword is used to begin a for loop.
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}
```
* <kbd> goto </kbd>
The `goto` keyword offers a jump to a labeled statement without any condition.
* <kbd> if </kbd>

The `if` statement is used to check a certain condition within a loop.
```go
package main

import "fmt"

func main() {
    x := 10

    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is not greater than 5")
    }
}
```
* <kbd> range </kbd>
The `range` keyword lets you iterate items over the list items like a map or an array.
* <kbd> return </kbd>
Go allows you to use the return values as variables and you can use the `return` keyword for that purpose.
* <kbd> select </kbd>
The `select` keyword lets a goroutine to wait during the simultaneous communication operations.
* <kbd> switch </kbd>

The `switch` statement is used to start a loop and use the if-else logic within the block.
```go
package main

import "fmt"

func main() {
    day := "Saturday"

    switch day {
    case "Monday":
        fmt.Println("It's the start of the week.")
    case "Saturday", "Sunday":
        fmt.Println("It's the weekend!")
    default:
        fmt.Println("It's a regular weekday.")
    }
}
```
### Function Modifier

* <kbd> defer </kbd>
The `defer` keyword is used to defer the execution of a function until the surrounding function executes.
* <kbd> go </kbd>
The `go` keyword triggers a goroutine which is managed by the golang-runtime.


