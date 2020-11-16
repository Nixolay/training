// Package selecting shows an example of balancing channels in select
package selecting

// GoSelect example of balancing channels in select.
func GoSelect() []int {
	const lenRange = 1000

	var a, b, c, d, e, f, g int

	ac := make(chan int)
	bc := make(chan int)
	cc := make(chan int)
	dc := make(chan int)
	ec := make(chan int)
	fc := make(chan int)
	gc := make(chan int)

	close(ac)
	close(bc)
	close(cc)
	close(dc)
	close(ec)
	close(fc)
	close(gc)

	for range [lenRange]int{} {
		select {
		case <-ac:
			a++
		case <-bc:
			b++
		case <-cc:
			c++
		case <-dc:
			d++
		case <-ec:
			e++
		case <-fc:
			f++
		case <-gc:
			g++
		}
	}

	return []int{a, b, c, d, e, f, g}
}
