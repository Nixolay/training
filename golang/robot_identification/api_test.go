package main_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	. "github.com/Nixolay/training/golang/robot_identification"
	"github.com/pkg/profile"
	. "github.com/smartystreets/goconvey/convey"
)

//nolint:wsl,scopelint
func TestUserCount(t *testing.T) {
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
		{name: "zero robots", timeout: time.Millisecond, sleep: time.Millisecond * 2},
		{name: "ten robots", timeout: time.Second, sleep: time.Second * 10, count: 10},
	}

	defer profile.Start(profile.ProfilePath("/tmp/profile")).Stop()

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

			time.Sleep(time.Millisecond * 2)

			So(uStorage.CountRobots(), ShouldEqual, d.count)
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
