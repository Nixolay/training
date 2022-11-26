package findseat_test

import (
	"fmt"
	"testing"

	findseat "github.com/Nixolay/training/codewar/golang/find_seat"
	"github.com/stretchr/testify/require"
)

/*
Места в кинотеатре расположены в один ряд. Только что пришедший зритель выбирает место, чтобы сидеть максимально далеко от остальных зрителей в ряду. То есть расстояние от того места, куда сядет зритель до ближайшего к нему зрителя должно быть максимально.
Гарантируется, что в ряду всегда есть свободные места и уже сидит хотя бы один зритель.
Напишите функцию, которая по заданному ряду мест (массиву из нулей и единиц) вернёт расстояние от выбранного пришедшим зрителем места до другого ближайшего зрителя.
*/

func TestFindSeat(t *testing.T) {
	data := []struct {
		seats []int
		fp    int
	}{
		{seats: []int{1, 0, 0, 0, 1}, fp: 1},
		{seats: []int{1, 0, 0, 0, 0}, fp: 3},
		{seats: []int{0, 0, 0, 0, 1, 0, 1}, fp: 3},
		{seats: []int{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, fp: 4},
		{seats: []int{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1}, fp: 4},
		{seats: []int{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1}, fp: 3},
		{seats: []int{0, 1, 0, 0, 0, 0, 1}, fp: 2},
		{seats: []int{1, 1, 1, 0, 1}, fp: 0},
		{seats: []int{1, 0, 1, 0, 0, 1, 0, 0, 0, 1}, fp: 1},
		{seats: []int{1, 0, 1, 0}, fp: 0},
		{seats: []int{1, 0, 0, 0, 0, 0, 0, 1, 0, 0}, fp: 3},
		{seats: []int{1, 0, 0, 1, 0, 0}, fp: 1},
		{seats: []int{1, 0, 0, 0, 1, 0, 0, 0}, fp: 2},
		{seats: []int{1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, fp: 3},
	}

	for i, td := range data {
		require.Equal(t, td.fp, findseat.FindSeat(td.seats), fmt.Sprint("test number: ", i))
	}
}
