package main_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/Nixolay/training/golang/slowly"
	. "github.com/smartystreets/goconvey/convey"
)

// nolint:gosec,noctx
func Test_slow(t *testing.T) {
	srv := httptest.NewServer(GetHandlers(5 * time.Second))
	defer srv.Close()

	const contentType = "application/json"

	var apiURL string = srv.URL + "/api/slow"

	Convey("Test API", t, func() {
		Convey("Checking url availability", func() {
			res, err := http.Post(apiURL, contentType, nil)
			if err := res.Body.Close(); err != nil {
				t.Fatal(err)
			}

			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldNotEqual, 404)
		})

		Convey("Invalid request type", func() {
			res, err := http.Get(apiURL)
			if err := res.Body.Close(); err != nil {
				t.Fatal(err)
			}

			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, 404)
		})

		Convey("Bad URL", func() {
			res, err := http.Get(fmt.Sprintf("%s/slow", srv.URL)) //nolint:noctx
			if err := res.Body.Close(); err != nil {
				t.Fatal(err)
			}

			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, 404)
		})
	})
}

// nolint:gosec,noctx
func TestTimouts_slow(t *testing.T) {
	srv := httptest.NewServer(GetHandlers(time.Second / 2))
	defer srv.Close()

	const contentType = "application/json"

	var apiURL string = srv.URL + "/api/slow"

	Convey("Test timeouts", t, func() {
		Convey("Send timeout", func() {
			buffer := bytes.NewBufferString(`{"timeout": 300}`)
			expected := `{"status":"ok"}`

			res, err := http.Post(apiURL, contentType, buffer)
			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(res.Header.Get("Content-Type"), ShouldEqual, "application/json")

			body, err := ioutil.ReadAll(res.Body)
			So(err, ShouldBeNil)

			defer func() {
				if err := res.Body.Close(); err != nil {
					t.Fatal(err)
				}
			}()

			So(string(body), ShouldEqual, expected)
		})

		Convey("Send timeout too long", func() {
			buffer := bytes.NewBufferString(fmt.Sprintf(
				"{\"timeout\": %d}",
				(time.Minute*5)/time.Millisecond))
			expected := "timeout too long"

			res, err := http.Post(apiURL, contentType, buffer)
			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, 503)

			defer func() {
				if err := res.Body.Close(); err != nil {
					t.Fatal(err)
				}
			}()

			body, err := ioutil.ReadAll(res.Body)
			So(err, ShouldBeNil)
			So(string(body), ShouldEqual, expected)
		})
	})
}
