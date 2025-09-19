package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Реализуйте HTTP-сервер на Go с одним эндпоинтом.
// Требования
// Сервер слушает адрес 127.0.0.1:8080.
// Эндпоинт: GET /ping.
// Поведение:
// - Если запрос без параметров → вернуть 200 OK, тело: строка "pong".
// - Если запрос имеет параметр ?slow=1 → симулировать задержку дольше 1 секунды и вернуть 504 Gateway Timeout.
// - Ответ "pong" должен приходить не позже чем через 1 секунду.
// Ограничения
// - Реализовать нужно только функцию: func StartServer() error
// - Нельзя изменять main и секцию import.
// - ельзя печатать ничего в stdout/stderr.
// - Для успешного прохождения проверок тело для 504 не проверяется, но статус-код должен быть именно 504.
func TestHTTPPingWithTimeout(t *testing.T) {
	mu := http.NewServeMux()

	mu.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		var slow time.Duration
		fmt.Sscan(r.URL.Query().Get("slow"), &slow)
		if slow > 0 {
			time.Sleep(slow * time.Second)
			w.WriteHeader(http.StatusGatewayTimeout)
			return
		}

		w.Write([]byte("pong"))
	})

	srv := http.Server{Addr: ":8080", Handler: mu}
	go srv.ListenAndServe()

	defer srv.Close()

	time.Sleep(time.Second)

	res, err := http.Get("http://127.0.0.1:8080/ping?slow=1")
	require.NoError(t, err)
	pong, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	println(string(pong))
}
