package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_slow(t *testing.T) {
	srv := httptest.NewServer(handlers())
	defer srv.Close()

	contentType := "application/json"
	apiURL := fmt.Sprintf("%s/api/slow", srv.URL)

	Convey("Test API", t, func() {
		Convey("Checking url availability", func() {
			res, err := http.Post(apiURL, contentType, nil)
			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldNotEqual, 404)
		})

		Convey("Invalid request type", func() {
			res, err := http.Get(apiURL)
			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, 404)
		})

		Convey("Bad URL", func() {
			res, err := http.Get(fmt.Sprintf("%s/slow", srv.URL))
			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, 404)
		})

		Convey("Send timeout", func() {
			buffer := bytes.NewBufferString(`{"timeout": 3000}`)
			expected := `{"status":"ok"}`

			res, err := http.Post(apiURL, contentType, buffer)
			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, 200)

			body, err := ioutil.ReadAll(res.Body)
			So(err, ShouldBeNil)

			defer res.Body.Close()

			So(string(body), ShouldEqual, expected)
		})

		Convey("Send timeout too long", func() {
			buffer := bytes.NewBufferString(fmt.Sprintf(
				"{\"timeout\": %d}",
				(time.Minute*6)/time.Millisecond))
			expected := `{"error":"timeout too long"}`

			res, err := http.Post(apiURL, contentType, buffer)
			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, 400)

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			So(err, ShouldBeNil)
			So(string(body), ShouldEqual, expected)
		})
	})
}
