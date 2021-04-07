// Package whoseport для поиска программы использующего порт.
package whoseport

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// WhosePort найдет программу, которая использует порт.
func WhosePort(port uint) string {
	// whoseport() { lsof -i ":$1" | grep LISTEN }
	const (
		binLsof = "/bin/lsof"
		binGrep = "/bin/grep"
	)

	//nolint:gosec
	lsof := exec.Command(binLsof, "-i", fmt.Sprintf(":%d", port))
	grep := exec.Command(binGrep, "LISTEN")

	var out bytes.Buffer
	grep.Stdout = &out
	grep.Stderr = os.Stderr

	var err error
	if grep.Stdin, err = lsof.StdoutPipe(); err != nil {
		return ""
	}

	logError(grep.Start())
	logError(lsof.Run())

	return out.String()
}

func logError(err error) {
	if err != nil {
		panic(err)
	}
}
