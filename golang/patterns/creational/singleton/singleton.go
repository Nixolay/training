// Package singleton example pattern.
package singleton

type single struct {
	count int
}

var instance *single //nolint:gochecknoglobals

// Increment count in singleton.
func (s *single) Increment() int {
	s.count++

	return s.count
}

// Incrementer exported type.
type Incrementer interface {
	Increment() int
}

// GetInstance returning instance single.
func GetInstance() Incrementer {
	if instance == nil {
		instance = new(single)
	}

	return instance
}
