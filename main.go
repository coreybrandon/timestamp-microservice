// main.go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	handler := http.HandlerFunc(TimeHandler)
	fmt.Printf("Listening on %s...", port)
	http.ListenAndServe(":"+port, handler)
}
