```go
package main

import "fmt"

func main() {
	defer fmt.Println("runs last")
	fmt.Println("runs first")
}
```

# `defer` keyword in Go

`defer` schedules a function call to run when the surrounding function returns. Deferred calls are commonly used for cleanup, such as closing files, unlocking mutexes, or releasing resources.

## Key points

- Deferred calls run in last-in, first-out (LIFO) order.
- Arguments are evaluated when the `defer` statement executes.
- Deferred functions run even when the function returns early.
- A deferred call runs before the surrounding function returns to its caller.

```go
func example() {
	defer fmt.Println("first deferred")
	defer fmt.Println("second deferred")

	fmt.Println("body")
}

// Output:
// body
// second deferred
// first deferred
```

Use `defer` immediately after acquiring a resource so cleanup is not accidentally omitted:

```go
file, err := os.Open("data.txt")
if err != nil {
	return err
}
defer file.Close()
```
