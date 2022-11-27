package singleton

import (
	"fmt"
	"sync"
	"testing"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstance(wg *sync.WaitGroup) *single {
	lock.Lock()
	defer lock.Unlock()
	defer wg.Done()

	if singleInstance == nil {
		fmt.Println("Создаем Singleton")
		singleInstance = &single{}
	} else {
		fmt.Println("Singleton уже создан")
	}

	return singleInstance
}

func TestSingleton(t *testing.T) {
	wg := &sync.WaitGroup{}

	wg.Add(30)
	for i := 0; i < 30; i++ {
		go getInstance(wg)
	}

	wg.Wait()
}
