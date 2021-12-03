// server.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Date struct {
	Unix *int64  `json:"unix"`
	Utc  *string `json:"utc"`
}

func TimeServer(w http.ResponseWriter, r *http.Request) {

	api := strings.TrimPrefix(r.URL.Path, "/api/")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("charset", "utf-8")

	t, err := getTimeStamp(api)
	if err != nil {
		if err := json.NewEncoder(w).Encode(Date{nil, nil}); err != nil {
			http.Error(w, "404", http.StatusInternalServerError)
		}
		return
	}

	dateString := t.Format("2006-01-02 15:04:05.999999999 GMT")
	timeString := t.Unix()

	date := Date{
		Unix: &timeString,
		Utc:  &dateString}

	if err := json.NewEncoder(w).Encode(&date); err != nil {
		http.Error(w, "Json string failed to marshal", http.StatusInternalServerError)
		return
	}

}

func getTimeStamp(t string) (time.Time, error) {

	if unixTime, err := strconv.Atoi(t); err == nil {
		return (time.Unix(int64(unixTime), 0)).UTC(), nil
	}

	parsedTime, err := time.Parse("2006-01-02", t)
	if err != nil {
		return time.Time{}, fmt.Errorf("cannot parse general time: %s", err)
	}

	return parsedTime.UTC(), nil
}
