package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

// Сделайте эндпоинт POST /order, который принимает JSON с полем request_id.
// Если приходит новый request_id, сервер возвращает 201 Created и сохраняет заказ.
// Если тот же request_id повторяется, сервер возвращает 200 OK без создания нового заказа.
// Sample Input:
// Sample Output: OK
func TestPostOrder(t *testing.T) {
	mu := http.NewServeMux()

	data := map[string]struct{}{}

	mu.HandleFunc("POST /order", func(w http.ResponseWriter, r *http.Request) {
		var req map[string]string
		json.NewDecoder(r.Body).Decode(req)

		id := req["request_id"]
		if _, ok := data[id]; ok {
			w.WriteHeader(http.StatusOK)
			return
		}

		data[id] = struct{}{}
		w.WriteHeader(http.StatusCreated)
	})

	srv := http.Server{Addr: ":8080", Handler: mu}
	defer srv.Close()
	go srv.ListenAndServe()

	time.Sleep(time.Second * 2)

	rsp, err := http.Post("http://localhost:8080/order", "", strings.NewReader(`{"request_id":"1234"}`))
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, rsp.StatusCode)

	rsp, err = http.Post("http://localhost:8080/order", "", strings.NewReader(`{"request_id":"1234"}`))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rsp.StatusCode)
}

// Реализуйте сервер с эндпоинтом POST /task, который принимает задачу и кладёт её в очередь (in-memory slice/chan).
// В ответ сразу возвращается 202 Accepted.
// Дополнительно сделайте эндпоинт GET /tasks, который показывает список всех ещё не выполненных задач.
func TestTasksQueue(t *testing.T) {
	mu := http.NewServeMux()

	tasks := []interface{}{}

	type Response struct {
		Data []interface{} `json:"data,omitempty"`
	}

	mu.HandleFunc("POST /task", func(w http.ResponseWriter, r *http.Request) {
		var req interface{}
		json.NewDecoder(r.Body).Decode(&req)
		tasks = append(tasks, req)
		w.WriteHeader(http.StatusAccepted)
	})
	mu.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Response{Data: tasks})
		w.Header().Set("Content-Type", "application/json")
	})

	srv := http.Server{Addr: ":8080", Handler: mu}
	defer srv.Close()
	go srv.ListenAndServe()
	time.Sleep(time.Second * 2)
}

// Сервер должен добавлять уникальный заголовок X-Request-ID ко всем ответам.
// Если клиент передал этот заголовок в запросе, сервер обязан вернуть его же.
func TestRequestID(t *testing.T) {
	mu := http.NewServeMux()

	XRequestID := 0

	mu.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")

		if id == "" {
			XRequestID++
			id = fmt.Sprint(XRequestID)
		}

		w.Header().Set("X-Request-ID", id)
	})

	srv := http.Server{Addr: ":8080", Handler: mu}
	defer srv.Close()
	go srv.ListenAndServe()

	time.Sleep(time.Second)

	client := http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/get", nil)
	rsp, err := client.Do(req)
	require.NoError(t, err)
	require.NotEmpty(t, rsp.Header.Get("X-Request-ID"))

	req.Header.Set("X-Request-ID", "10")
	rsp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, "10", rsp.Header.Get("X-Request-ID"))
}
