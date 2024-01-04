## io.Reader and io.Writer

io.Reader Interface:
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```