# learn-go
Trong Go, Chỉ có public, private 
- 



### Map
```go



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
* const
The `const` keyword is used to introduce a name for a scalar value like 2 or 3.14, etc.
* var
The `var` keyword is used to create the variables in the `Go` language.
* func
The `func` keyword is used to declare a function.
* type
We can use the `type` keyword to introduce a new struct type.
* import
The `import` keyword is used to import packages.
* package
The code is grouped as a unit in a package. The `package` keyword is used to define one.
### Composite Types
* chan
The `chan` keyword is used to define a channel. In `Go`, you are allowed to run parallel pieces of code simultaneously.
* interface
The `interface` keyword is used to specify a method set. A method set is a list of methods for a type.
* map
The `map` keyword defines a map type. A map os an unordered collection of the key-value pairs.
* struct
Struct is a collection of fields. We can use the `struct` keyword followed by the field declarations.

### Control Flow

* break
The `break` keyword lets you terminate a loop and continue execution of the rest of the code.
* case 
This is a form of a `switch` construct. We specify a variable after the switch.
* continue
The `continue` keyword allows you to go back to the beginning of the ‘for’ loop.
* default
The `default` statement is optional but you need to use either case or default within the switch statement. The switch jumps to the default value if the case does not match the controlling expression.
* else
The `else` keyword is used with the `if` keyword to execute an alternate statement if a certain condition is false.
* fallthrough
This keyword is used in a case within a switch statement. When we use this keyword, the next case is entered.
* for
The `for` keyword is used to begin a for loop.
* goto
The `goto` keyword offers a jump to a labeled statement without any condition.
* if
The `if` statement is used to check a certain condition within a loop.
* range
The `range` keyword lets you iterate items over the list items like a map or an array.
* return
Go allows you to use the return values as variables and you can use the `return` keyword for that purpose.
* select
The `select` keyword lets a goroutine to wait during the simultaneous communication operations.
* switch
The `switch` statement is used to start a loop and use the if-else logic within the block.
### Function Modifier

* defer
The `defer` keyword is used to defer the execution of a function until the surrounding function executes.
go
The `go` keyword triggers a goroutine which is managed by the golang-runtime.


