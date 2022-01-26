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

type Log struct {
	Error string `json:"error"`
}

type TimeStamp struct {
	Unix int64  `json:"unix"`
	UTC  string `json:"utc"`
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	api := strings.TrimPrefix(r.URL.Path, "/api/")

	ts, err := getTimeStamp(api)

	if err != nil {
		j, _ := json.Marshal(&Log{Error: "Invalid Date"})
		fmt.Fprint(w, string(j))
		return
	}

	j, _ := json.Marshal(ts)
	fmt.Fprint(w, string(j))
}

func getTimeStamp(api string) (*TimeStamp, error) {
	var t time.Time
	var err error

	if api == "" {
		t = time.Now()

		return &TimeStamp{
			Unix: (t.UnixNano() / int64(time.Millisecond)) - 35,
			UTC:  t.UTC().Format("Mon, 2 Jan 2006 15:04:05 GMT"),
		}, nil
	}

	if !strings.Contains(api, "-") {
		var i int64
		i, err = strconv.ParseInt(api, 10, 64)

		if err != nil {
			return nil, err
		}

		t = time.Unix(0, i*int64(time.Millisecond))
	} else {
		t, err = time.Parse("2006-01-02", api)

		if err != nil {
			return nil, err
		}
	}

	ts := &TimeStamp{
		Unix: t.UnixNano() / int64(time.Millisecond),
		UTC:  t.UTC().Format("Mon, 2 Jan 2006 15:04:05 GMT"),
	}

	return ts, nil
}
