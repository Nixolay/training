package selecting_test

import (
	"fmt"
	"sync"
	"testing"

	. "github.com/Nixolay/training/golang/selecting"
)

func TestGoSelecting(t *testing.T) {
	wg := sync.WaitGroup{}

	for range [1000]int{} {
		wg.Add(1)

		go func() {
			defer wg.Done()

			data := GoSelect()
			for _, item := range data {
				if item < 85 {
					panic(fmt.Sprintf("Item: %d, selection is bad!", item))
				}
			}
		}()
	}

	wg.Wait()
}
