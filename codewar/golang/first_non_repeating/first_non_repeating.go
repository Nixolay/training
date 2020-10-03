package repair

import (
	"strings"
)

type Data struct {
	r               rune
	position, count int
}

func FirstNonRepeating(str string) string {
	data := make(map[string]*Data)

	for iterator, r := range str {
		d, ok := data[strings.ToLower(string(r))]
		if !ok {
			data[strings.ToLower(string(r))] = &Data{r: r, position: iterator, count: 1}
		}

		if ok {
			d.count++
		}
	}

	position := Data{position: -1}

	for _, d := range data {
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
