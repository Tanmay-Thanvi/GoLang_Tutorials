WaitGroups in Go, available in the `sync` package, are a powerful tool for synchronizing goroutines. They're used to wait for a collection of goroutines to finish executing their tasks before proceeding further in the main goroutine or another function.

### Basics of WaitGroups:

1. **Add():** Before starting a goroutine, you increment the WaitGroup counter using `Add()`.

2. **Done():** At the end of each goroutine's task, you call `Done()` to decrement the WaitGroup counter.

3. **Wait():** Finally, you use `Wait()` to block the main goroutine until the WaitGroup counter becomes zero, indicating that all the goroutines have completed their tasks.

### Example:

```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	fmt.Printf("Worker %d starting\n", id)
	// Simulate some work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the counter for each goroutine
		go worker(i, &wg)
	}

	wg.Wait() // Block until all the goroutines finish
	fmt.Println("All workers have completed")
}
```

### Key Points:

- Use `Add(n)` to add `n` to the WaitGroup counter before starting `n` goroutines.
- Each goroutine must call `Done()` when it completes its task to decrement the counter.
- `Wait()` blocks the main goroutine until the counter becomes zero.

### Common Use Cases:

- Waiting for a batch of parallel tasks to complete before proceeding.
- Synchronization between a fixed number of goroutines.
- Managing concurrency in tasks like waiting for a group of HTTP requests to finish.

### Thread Safety:

WaitGroups are safe for concurrent use. Multiple goroutines can safely increment and decrement the counter concurrently without race conditions.

### Best Practices:

- Always call `Done()` after the goroutine completes, even in error scenarios.
- Avoid manipulating the counter directly; always use `Add()` and `Done()` in pairs.
- Ensure that every `Add()` call has a corresponding `Done()` call to prevent deadlocks.

WaitGroups are an essential tool for coordinating concurrent tasks in Go, enabling effective synchronization among goroutines and allowing clean and controlled exits from concurrent operations.