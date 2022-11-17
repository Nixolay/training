package deduplicate

func Deduplicate(data []int) []int {
	arr := make([]int, 0, len(data))

	for i := range data {
		if i != 0 && data[i-1] == data[i] {
			continue
		}

		arr = append(arr, data[i])
	}

	return arr
}

type List struct {
	Value int
	Next  *List
}

func (l List) Slice() []int {
	var s []int
	for e := &l; e != nil; e = e.Next {
		s = append(s, e.Value)
	}

	return s
}

func DeduplicateList(list *List) {
	keys := make(map[int]struct{})

	for element := list; element.Next != nil; {
		if _, ok := keys[element.Next.Value]; ok || element.Value == element.Next.Value {
			element.Next = element.Next.Next

			continue
		}

		keys[element.Value] = struct{}{}
		element = element.Next
	}
}
