/*Package repair for non repeating.*/
package repair

import (
	"strings"
)

type data struct {
	r               rune
	position, count int
}

// FirstNonRepeating non repeat.
func FirstNonRepeating(str string) string {
	dt := make(map[string]*data)

	for iterator, r := range str {
		d, ok := dt[strings.ToLower(string(r))]
		if !ok {
			dt[strings.ToLower(string(r))] = &data{r: r, position: iterator, count: 1}
		}

		if ok {
			d.count++
		}
	}

	position := data{position: -1}

	for _, d := range dt {
		if d.count > 1 {
			continue
		}

		if position.position == -1 {
			position = *d

			continue
		}

		if position.position > d.position {
			position = *d
		}
	}

	if position.position == -1 {
		return ""
	}

	return string(position.r)
}
