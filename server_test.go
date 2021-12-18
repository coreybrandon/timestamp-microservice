//server_test.go

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

// Didn't end up using this.

// type StubTimeStamp struct {
// 	Unix int64  `json:"unix"`
// 	UTC  string `json:"utc"`
// }

func TestGetTimeStamp(t *testing.T) {
	t.Run("returns Unix Timestamp", func(t *testing.T) {
		request := newGetTimeStampRequest("1451001600000")
		response := httptest.NewRecorder()

		TimeHandler(response, request)

		assertResponseBody(t, response.Body.String(), "{\"unix\":1451001600000,\"utc\":\"Fri, 25 Dec 2015 00:00:00 GMT\"}")
	})

	t.Run("returns UTC Timestamp", func(t *testing.T) {
		request := newGetTimeStampRequest("2015-12-25")
		response := httptest.NewRecorder()

		TimeHandler(response, request)

		assertResponseBody(t, response.Body.String(), "{\"unix\":1451001600000,\"utc\":\"Fri, 25 Dec 2015 00:00:00 GMT\"}")
	})

	t.Run("returns { error : 'Invalid Date' } if the input string is invalid", func(t *testing.T) {
		request := newGetTimeStampRequest("3-26=36=36")
		response := httptest.NewRecorder()

		TimeHandler(response, request)

		assertResponseBody(t, response.Body.String(), "{\"error\":\"Invalid Date\"}")
	})

	t.Run("empty string returns current time JSON with unix + utc keys", func(t *testing.T) {
		request, _ := getTimeStamp("")

		utcNow := time.Now().UTC().Format("Mon, 2 Jan 2006 15:04:05 GMT")
		unixNow := time.Now().Unix()

		if !strings.Contains(request.UTC, utcNow[:20]) {
			t.Error("does not return UTC Timestamp")
		}

		if !strings.Contains(strconv.FormatInt(request.Unix, 10), strconv.FormatInt(unixNow, 10)[:5]) {
			t.Error("does not return Unix Timestamp")
		}
	})
}

func newGetTimeStampRequest(date string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/%s", date), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
