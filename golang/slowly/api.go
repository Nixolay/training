package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	timeoutSeconds = 5
	pathSlow       = "/api/slow"
)

type message struct {
	Timeout time.Duration `json:"timeout,omitempty"`
	Status  string        `json:"status,omitempty"`
	Error   string        `json:"error,omitempty"`
}

// RunServer starts the server.
func RunServer(addr string) error {
	return http.ListenAndServe(addr, GetHandlers())
}

// GetHandlers returns prepared handlers.
func GetHandlers() http.Handler {
	r := http.NewServeMux()

	slowHandler := http.HandlerFunc(slow)
	timeoutHandler := http.TimeoutHandler(
		slowHandler, timeoutSeconds*time.Second,
		"timeout too long",
	)

	r.Handle(pathSlow, timeoutHandler)

	return r
}

// slow handle request.
func slow(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	// checked the correct method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)

		if _, err := w.Write([]byte(http.StatusText(http.StatusNotFound))); err != nil {
			log.Println(err)
		}

		return
	}

	// reading request body
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	// reading message
	msg := message{}
	if err := json.Unmarshal(content, &msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	time.Sleep(msg.Timeout * time.Millisecond)

	bodyResponse, err := json.Marshal(message{Status: "ok"}) //nolint:exhaustivestruct
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(bodyResponse); err != nil {
		log.Println(err)
	}
}
