package bubble_test

import (
	"reflect"
	"testing"

	"github.com/Nixolay/training/golang/sorts/bubble"
)

func TestBubbleSort(t *testing.T) {
	actual := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	expected := []int{2, 14, 33, 37, 45, 45, 45, 91, 99, 102, 102, 104, 106, 106, 109, 212, 501, 3001, 7800, 9932}

	bubble.Sort(actual)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("массивы не равны")
	}
}
