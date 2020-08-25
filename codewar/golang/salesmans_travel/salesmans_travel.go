package salesmans

import "strings"

func Travel(r, zipcode string) string {
	if zipcode == "" || r == "" {
		return ":/"
	}

	out := zipcode + ":"
	end := ""

	for _, v := range strings.Split(r, ",") {
		if v[len(v)-len(zipcode):] == zipcode {
			v = strings.TrimSpace(v)
			i := strings.Index(v, " ")
			out += v[i+1:len(v)-len(zipcode)-1] + ","
			end += "," + v[:i]
		}
	}

	if end == "" {
		return zipcode + ":/"
	}

	return out[:len(out)-1] + "/" + end[1:]
}
