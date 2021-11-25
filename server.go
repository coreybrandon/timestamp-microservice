// server.go
package main

import (
	"fmt"
	"net/http"
)

// Using fmt.Fprint to send strings as HTTP responses
func TimeServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "1451001600000")
}
