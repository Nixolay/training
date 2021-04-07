package whoseport_test

import (
	"testing"

	"github.com/Nixolay/training/golang/system_utility/whoseport"
)

func TestWhosePort(t *testing.T) {
	const port = 8080

	whoseport.WhosePort(port)
}
