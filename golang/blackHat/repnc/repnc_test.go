package repnc

import (
	"io"
	"net"
	"os/exec"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	t.Skip()

	listener, err := net.Listen("tcp", ":8080")
	require.NoError(t, err, "Слушаем порт")

	for {
		conn, err := listener.Accept()
		require.NoError(t, err, "Ожидание сообщений")

		go Handle(t, conn)
	}
}

// telnet localhost 8080.
func Handle(t *testing.T, conn net.Conn) {
	// после того как отработает наша функция, соединение нужно закрыть
	defer conn.Close()

	cmd := exec.Command("/bin/sh", "-i")
	if runtime.GOOS == "windows" {
		// в windows другие пути и обычно не установлен sh но есть cmd
		cmd = exec.Command("cmd.exe")
	}

	// Из-за некоторой специфики windows, мы не можем просто передать в stdout наш conn
	// нам нужно принудительно очищать stdout, с этим прекрасно справится io.Pipe.
	pReader, pWriter := io.Pipe()

	// Stdin io.Reader, в conn есть метод Reader, а значит удовлетворяет интерфейсу io.Reader
	cmd.Stdin = conn
	// Stdout io.Writer, pWriter удовлетворяет интерфейсу io.Writer
	cmd.Stdout = pWriter

	go func() {
		// io.Copy() копирует из src в dst до тех пор, пока в src не будет
		// достигнут EOF или не возникнет ошибка. Он возвращает количество
		// скопированных байт и первую ошибку, возникшую при копировании,
		// если таковая имеется.
		_, err := io.Copy(conn, pReader)
		require.NoError(t, err, "попытка прочитать сообщение в pReader")
	}()

	require.NoError(t, cmd.Run(), "запускаем программу - windows: cmd | others: sh")
}
