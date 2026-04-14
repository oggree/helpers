# Oggree Helpers

A lightweight utility package providing commonly used helper functions for the Oggree project. 
It uses modern Go 1.18+ Generics for type-sensitive array lookups and utilizes Go 1.22's `math/rand/v2` for secure, deterministic pseudorandom generation.

## Installation

To add this library to your Go project:

```bash
go get github.com/oggree/helpers
```

## Functions

### `InArray`
Checks if a specific value exists within a slice of comparable items. It returns `true` and the first index if found, or `false` and `-1` if not found.

```go
package main

import (
	"fmt"
	"github.com/oggree/helpers"
)

func main() {
	fruits := []string{"apple", "banana", "cherry"}
	
	exists, index := helpers.InArray("banana", fruits)
	if exists {
		fmt.Printf("Found at index: %d\n", index) // Outputs: Found at index: 1
	}
}
```

### `RandomInt`
Generates a pseudo-random integer in the half-open interval `[min, max)`.
*Note: Panics if `min >= max`, consistent with the standard library's behavior.*

```go
package main

import (
	"fmt"
	"github.com/oggree/helpers"
)

func main() {
	// Generates a random integer from 5 up to but not including 10
	val := helpers.RandomInt(5, 10)
	fmt.Println(val)
}
```

### `RandomString`
Generates a random string of a specified length consisting solely of uppercase A-Z characters.

```go
package main

import (
	"fmt"
	"github.com/oggree/helpers"
)

func main() {
	// Generates a 16-character completely random string (e.g. "QKXWPLMZV...")
	token := helpers.RandomString(16)
	fmt.Println(token)
}
```

## Testing
To run the included test suite over this library:

```bash
go test -v ./...
```
