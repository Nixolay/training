package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" //nolint:gci,gosec
	"time"

	"golang.org/x/net/http2"
)

var userStorage = CreateUserStorage(time.Minute) //nolint:gochecknoglobals

// RunServer Server API launch.
func RunServer() {
	httpServer := http.Server{Addr: ":8080"} //nolint:exhaustivestruct

	http2Server := http2.Server{}
	if err := http2.ConfigureServer(&httpServer, &http2Server); err != nil {
		log.Fatal(err)
	}

	createHandlers()

	log.Fatal(httpServer.ListenAndServe())
}

func createHandlers() {
	http.HandleFunc("/count", count)
	http.HandleFunc("/", identification)
}

func count(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	fmt.Fprint(w, userStorage.CountRobots())
}

func identification(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if keys, ok := r.URL.Query()["user_id"]; ok && len(keys) > 0 {
		userStorage.Inc(keys[0])
	}

	fmt.Fprint(w, "ok")
}
