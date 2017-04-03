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
		fmt.Println(i)
	}
}
```

## Efficiency Concerns

This package was created specifically to address the critiques of https://github.com/bradfitz/iter found in this [StackOverflow discussion](http://stackoverflow.com/questions/21950244/is-there-a-way-to-iterate-over-a-range-of-integers-in-golang) while providing the same functionality.

As can be seen in its simple implementation, it is backed by a simple idiomatic `for` loop, creating no new data structures aside from a `chan int` over which to send values. This should minimize cache issues, but at the cost of speed.