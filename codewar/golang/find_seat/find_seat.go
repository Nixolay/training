package findseat

type Begin struct {
	count int
	ok    bool
}

func FindSeat(row []int) int {
	var (
		maxAway, count int
		begin          = Begin{ok: true}
	)

	for i, seat := range row {
		if i == 0 && seat == 1 {
			begin.ok = false
		}

		if seat == 1 && maxAway < count {
			maxAway = count

			if begin.ok && begin.count == 0 {
				begin.count = count
			}
		}

		if seat == 1 {
			count = 0

			continue
		}

		count++
	}

	switch {
	case maxAway == 1:
		return 0
	case count > 0 && count-1 >= maxAway/2:
		return count - 1
	case begin.ok && begin.count-1 >= maxAway/2:
		return begin.count - 1
	default:
		return maxAway / 2
	}
}
