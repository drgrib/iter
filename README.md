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

As discussed in this [StackOverflow answer](http://stackoverflow.com/a/43192236/130427), it ultimately produced a very expensive option for iteration. Within its benchmark, it also shows how the `bradfitz` option underperforms the built-in `for` loop. Ultimately, if maximum efficiency is your goal, the built-in `for` clause is the best option. And as discussed in my answer, you might not miss the Python `range()` approach as much as you think you will before writing a lot of Go code. 