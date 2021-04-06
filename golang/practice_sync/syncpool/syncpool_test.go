package syncpool_test

import (
	"testing"

	"github.com/Nixolay/training/golang/practice_sync/syncpool"
)

func BenchmarkPool(b *testing.B) {
	syncpool.T()
}
