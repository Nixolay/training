package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const pathSlow = "/api/slow"

type message struct {
	Timeout time.Duration `json:"timeout,omitempty"`
	Status  string        `json:"status,omitempty"`
	Error   string        `json:"error,omitempty"`
}

func RunServer(addr string) error {
	return http.ListenAndServe(addr, handlers())
}

func handlers() http.Handler {
	r := http.NewServeMux()

	slowHandler := http.HandlerFunc(slow)
	r.Handle(pathSlow, middlewareSlow(slowHandler))

	return r
}

func slow(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	msg := message{}
	if err := json.Unmarshal(content, &msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	time.Sleep(msg.Timeout * time.Millisecond)

	sendRequest(w, message{Status: "ok"}, http.StatusOK)
}

func middlewareSlow(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == pathSlow {
			go next.ServeHTTP(w, r)

			time.Sleep(5 * time.Second)
			sendRequest(w, message{Error: "timeout too long"}, http.StatusBadRequest)
		}
	})
}

func sendRequest(w http.ResponseWriter, msg message, status int) {
	if w.Header().Get("Content-Type") != "" {
		return
	}

	body, err := json.Marshal(msg)
	if err != nil {
		log.Print(err)
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(body); err != nil {
		log.Print(err)
	}
}
