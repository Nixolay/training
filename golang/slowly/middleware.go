package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

const timeoutSeconds = 5

// middlewareSlow if the request takes a long time to process, it sends an error.
func middlewareSlow(next http.Handler) http.Handler {
	return &timeoutSlowHandler{handler: next}
}

type timeoutSlowHandler struct {
	handler     http.Handler
	testContext context.Context
}

func (h *timeoutSlowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := h.testContext
	dt := time.Second * timeoutSeconds

	if ctx == nil {
		var cancelCtx context.CancelFunc
		ctx, cancelCtx = context.WithTimeout(r.Context(), dt)

		defer cancelCtx()
	}

	r = r.WithContext(ctx)

	done := make(chan struct{})

	panicChan := make(chan interface{}, 1)

	tw := &slowWriter{
		w: w,
	}

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		h.handler.ServeHTTP(tw, r)

		close(done)
	}()

	select {
	case <-done:
	case err := <-panicChan:
		log.Println(err)
	case <-ctx.Done():
		tw.WriteHeader(http.StatusBadRequest)

		body, err := json.Marshal(message{Error: "timeout too long"})
		if err != nil {
			log.Print(err)
		}

		_, err = tw.Write(body)
		if err != nil {
			log.Println(err)
		}
	}
}

type slowWriter struct {
	w http.ResponseWriter
	h http.Header

	mu          sync.RWMutex
	wroteHeader bool
}

func (tw *slowWriter) Header() http.Header { return tw.h }

func (tw *slowWriter) Write(p []byte) (n int, err error) {
	tw.mu.Lock()
	defer tw.mu.Unlock()

	return tw.w.Write(p)
}

func (tw *slowWriter) WriteHeader(code int) {
	tw.mu.Lock()
	defer tw.mu.Unlock()

	tw.wroteHeader = true

	tw.w.Header().Set("Content-Type", "application/json")
	tw.w.WriteHeader(code)
}
