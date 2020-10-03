/*
Complete the solution so that it splits the string into pairs of two characters.
If the string contains an odd number of characters then it should replace the
missing second character of the final pair with an underscore ('_').

Examples:
 Solution("abc") //should return ["ab", "c_"]
 Solution("abcdef") //should return ["ab", "cd", "ef"]
*/
package splitstrings

import "regexp"

func Solution(str string) []string {
	if len(str)%2 != 0 {
		str += "_"
	}

	two := 2
	out := make([]string, 0, len(str)/two)

	for i := 0; i < len(str)/2+2; i += 2 {
		out = append(out, str[i:i+2])
	}

	return out
}

func Solution2(str string) []string {
	return regexp.MustCompile(".{2}").FindAllString(str+"_", -1)
}
