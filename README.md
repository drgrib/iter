This package is meant to provide a Go equivalent to the basic Python `range()` functionality found throughout idiomatic Python.

It can be executed as follows

```
package main

import (
	"fmt"
	"github.com/drgrib/iter"
)

func main() {
	for i := range iter.N(10) {
		fmt.Print(i)
	}
}
```
