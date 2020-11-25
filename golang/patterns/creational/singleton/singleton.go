// Package singleton example pattern.
package singleton

type single struct {
	count int
}

var instance *single //nolint:gochecknoglobals

// GetInstance returning instance single.
//nolint:golint
func GetInstance() *single {
	if instance == nil {
		instance = new(single)
	}

	return instance
}

// Increment count in singleton.
func (s *single) Increment() int {
	s.count++

	return s.count
}
