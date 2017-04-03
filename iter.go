// Package iter provides a maximally efficient and concise range over integers [0,n) backed internally by a simple for loop.
package iter

/*
N ranges over [0,n).

For example:

	for i := range iter.N(10) {
		fmt.Println(i)
	}

will print

	0
	1
	2
	3
	4
	5
	6
	7
	8
	9
*/
func N(n int) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
