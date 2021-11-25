package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Testing Unix Timestamp response
func TestGETUnix(t *testing.T) {
	t.Run("Returns a unix key that is a Unix timestamp of the input date in milliseconds", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "api/2015-12-25", nil) // Leveraging ResponseWriter to inspect what is being written by the handler
		response := httptest.NewRecorder()                                   // Inspecting what was written as a response

		TimeServer(response, request)

		got := response.Body.String()
		want := "1451001600000"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

// func TestGETUtc(t *testing.T) {
// 	t.Run("Returns a utc key that is a string of the input date", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/api/1451001600000", nil) // Leveraging ResponseWriter to inspect what is being written by the handler
// 		response := httptest.NewRecorder()                                       // Inspecting what was written as a response

// 		TimeServer(response, request)

// 		got := response.Body.String()
// 		want := "Fri, 25 Dec 2015 00:00:00 GMT"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// }
