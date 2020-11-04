// Package hidden for testing hidden parameters
package hidden

// IsHidden const data for testing.
const IsHidden = "is hidden"

// Hidden simple struct for testing hidden parameters.
type Hidden struct {
	DataNoHidden int
	h            string
}

// CreateHidden with base data for testing.
func CreateHidden() *Hidden {
	h := &Hidden{DataNoHidden: 1, h: IsHidden}

	return h
}
