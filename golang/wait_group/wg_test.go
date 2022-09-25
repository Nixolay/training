package waitgroup_test

import (
	"sync"
	"testing"
)

func TestWG(t *testing.T) {
	wg := WaitGroup{}
	wg.Add(1)
}

type WaitGrouper interface {
	Add(int)
	Done()
	Wait()
}

type WaitGroup struct {
	delta int
	ch    chan struct{}
	m     sync.Mutex
}

func (wg *WaitGroup) Add(delta int) {
	wg.m.Lock()
	defer wg.m.Unlock()
	wg.delta += delta
}

func (wg *WaitGroup) Done() {
	wg.m.Lock()
	defer wg.m.Unlock()

	wg.delta--

	if wg.delta < 0 {
		panic("Слишком много Done")
	}
	if wg.delta == 0 && wg.ch != nil {
		close(wg.ch)
	}
}

func (wg *WaitGroup) Wait() {
	if wg.delta == 0 {
		return
	}

	wg.ch = make(chan struct{})
	<-wg.ch
}
