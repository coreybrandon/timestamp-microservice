package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(TimeServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
