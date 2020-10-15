/*Package aretheythesame check deep equal.*/
package aretheythesame

import (
	"reflect"
	"sort"
)

// Comp check deep equal.
func Comp(a, b []int) bool {
	if a == nil || b == nil {
		return false
	}

	c, d := a, b
	for i, n := range a {
		c[i] = n * n
	}

	sort.Ints(c)
	sort.Ints(d)

	return reflect.DeepEqual(c, d)
}
