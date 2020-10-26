package main_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	. "github.com/Nixolay/training/golang/robot_identification"
	. "github.com/smartystreets/goconvey/convey"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)

	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)

	fmt.Println("----")
}

//nolint:wsl,scopelint
func TestUserCount(t *testing.T) {
	var mem runtime.MemStats
	var data []struct {
		name           string
		timeout, sleep time.Duration
		count          int
	} = []struct {
		name    string
		timeout time.Duration
		sleep   time.Duration
		count   int
	}{
		{name: "zero robots", timeout: 0, sleep: time.Millisecond},
		{name: "ten robots", timeout: time.Second, sleep: time.Second / 10, count: 10},
	}

	// defer profile.Start(profile.ProfilePath("/tmp/profile")).Stop()
	// defer profile.Start(profile.ProfilePath("/tmp/profile"), profile.TraceProfile).Stop()

	for _, d := range data {
		Convey(d.name, t, func() {
			uStorage := CreateUserStorage(d.timeout)

			wg := sync.WaitGroup{}

			for i := range [1000]int{} {
				wg.Add(1)

				go func(i int) {
					userRequests := 10

					if i%100 == 0 {
						userRequests = 150
					}

					for range make([]int, userRequests) {
						uStorage.Inc(fmt.Sprintf("User_%d", i))
					}

					wg.Done()
				}(i)
			}

			wg.Wait()

			time.Sleep(d.sleep)

			So(uStorage.CountRobots(), ShouldEqual, d.count)

			printStats(mem)
		})
	}
}

//nolint:wsl
func BenchmarkCountUsers(b *testing.B) {
	uStorage := CreateUserStorage(0)

	for i := range [1000]int{} {
		uStorage.Inc(fmt.Sprintf("User_%d", i))
	}

	uStorage.Inc("User_1")
	if uStorage.CountRobots() != 0 {
		panic("Incorrect count robots")
	}
}
