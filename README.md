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

## Efficiency Concerns

This package was created specifically to address the critiques for https://github.com/bradfitz/iter found in this [StackOverflow discussion](http://stackoverflow.com/questions/21950244/is-there-a-way-to-iterate-over-a-range-of-integers-in-golang).