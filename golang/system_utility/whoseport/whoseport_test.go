package whoseport_test

import (
	"testing"

	"github.com/Nixolay/training/golang/system_utility/whoseport"
)

func TestWhosePort(t *testing.T) {
	t.Skip("TestWhosePort skip in short mod")

	const port = 3000

	if process := whoseport.WhosePort(port); process == "" {
		t.Fatal("process not find")
	}
}
