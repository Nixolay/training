package blackproxy_test

import (
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

const dstURL = "golang.org:80"

func Handle(t *testing.T, src net.Conn) {
	dst, err := net.Dial("tcp", dstURL)
	require.NoError(t, err, "попытка подключения к: "+dstURL)

	go func() {
		_, err := io.Copy(dst, src)
		require.NoError(t, err, "попытка отправить сообщение на "+dstURL)
	}()

	_, err = io.Copy(src, dst)
	require.NoError(t, err, "попытка получить сообщение от "+dstURL)
}

func TestHandler(t *testing.T) {
	listener, err := net.Listen("tcp", ":8080")
	require.NoError(t, err, "Слушаем порт")

	for {
		conn, err := listener.Accept()
		require.NoError(t, err, "Ожидание сообщений")
		go Handle(t, conn)
	}
}
