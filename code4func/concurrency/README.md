## Concurrency

Concurrency is the composition of independent activities, e.g: deal with multiple user requests simultaneously or processing data in batches.

Go has two styles for writing programs:

- `traditional`: using threads
- `Go approach`: values as passed between independent activities (goroutines) to communicate processes

### Goroutines

A goroutine is a concurrent activity in a lightweight thread (not the traditional one you have in an operating system).

```go
func main() {
    login()
	go launch() // goroutine
}
```

or using an `anonymous` function:

```go
func main() {
    login()
	go func(){ // goroutine in an anonymous function
		launch()
}
}
```

### Channels
A channel in Go is a communication mechanism between goroutines. The idea is to send a value from one goroutine to another, using channels.

