//server_test.go

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTimeStamp(t *testing.T) {
	t.Run("returns Unix Timestamp", func(t *testing.T) {
		request := newGetTimeStampRequest("1451001600000")
		response := httptest.NewRecorder()

		TimeServer(response, request)

		assertResponseBody(t, response.Body.String(), "{\"unix\":1451001600000,\"utc\":\"Fri, 25 Dec 2015 00:00:00 GMT\"}")
	})

	// @Todo - Return the actual UTC Timestamp. So far I'm just getting Unix time back.

	t.Run("returns UTC Timestamp", func(t *testing.T) {
		request := newGetTimeStampRequest("2015-12-25")
		response := httptest.NewRecorder()

		TimeServer(response, request)

		assertResponseBody(t, response.Body.String(), "{\"unix\":1451001600000,\"utc\":\"Fri, 25 Dec 2015 00:00:00 GMT\"}")
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
