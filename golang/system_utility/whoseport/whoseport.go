// Package whoseport для поиска программы использующего порт.
package whoseport

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// PIDUsePort структура для получения необходимых данных.
type PIDUsePort struct {
	PID, PPID uint
	Value     string
}

// WhosePort функция для получения данных о программе котора использует порт.
func WhosePort(port uint) PIDUsePort {
	var whosePort PIDUsePort

	whosePort.Value = findProccess(getProcesses(), port)
	whosePort.PID, whosePort.PPID = getPIDAndPPID(whosePort.Value)

	return whosePort
}

func findProccess(processes []string, port uint) string {
	portString := fmt.Sprintf(":%d", port)
	println(portString)

	for _, process := range processes {
		if strings.Contains(process, portString) {
			return process
		}
	}

	return ""
}

// TODO: нужна реализация после того как найду нормальный скрипт поиска программы.
//nolint
func getPIDAndPPID(process string) (uint, uint) {
	return 0, 0
}

func getProcesses() []string {
	switch runtime.GOOS {
	case `linux`:
		return getLinuxProcesses()
	case `windows`:
	case `darwin`:
	}

	return nil
}

func getLinuxProcesses() []string {
	// netstat –lnt --program
	app := "/bin/netstat"
	cmd := exec.Command(app, "-tunlp", "--program")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil
	}

	return strings.Split(out.String(), "\n")
}
