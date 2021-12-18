# timestamp-microservice

### Description:

- An empty string should return the **current** time in a JSON object
- A request to `/api/` with a valid date should return a JSON object with a **unix** key that is a Unix timestamp of the input date in milliseconds
- A request to `/api/` with a valid date should return a JSON object with a **utc** key that is a string of the input date in the format: `Mon, 2 Jan 2006 15:04:05 GMT`
- A request to `/api/1451001600000` should return `{ unix: 1451001600000, utc: "Fri, 25 Dec 2015 00:00:00 GMT" }`
- If the input date string is invalid, the api returns an object having the structure `{ error : "Invalid Date" }`

Example usage:

```
https://localhost:5000/api/2015-12-25
https://localhost:5000/api/1451001600000
```

Example output:

```
{"unix":1451001600000, "utc":"Fri, 25 Dec 2015 00:00:00 GMT"}
```

### Test

```
go test -v
```

```
=== RUN   TestGetTimeStamp
=== RUN   TestGetTimeStamp/returns_Unix_Timestamp
=== RUN   TestGetTimeStamp/returns_UTC_Timestamp
=== RUN   TestGetTimeStamp/returns_{_error_:_'Invalid_Date'_}_if_the_input_string_is_invalid
=== RUN   TestGetTimeStamp/empty_string_returns_current_time_JSON_with_unix_+_utc_keys
--- PASS: TestGetTimeStamp (0.00s)
    --- PASS: TestGetTimeStamp/returns_Unix_Timestamp (0.00s)
    --- PASS: TestGetTimeStamp/returns_UTC_Timestamp (0.00s)
    --- PASS: TestGetTimeStamp/returns_{_error_:_'Invalid_Date'_}_if_the_input_string_is_invalid (0.00s)
    --- PASS: TestGetTimeStamp/empty_string_returns_current_time_JSON_with_unix_+_utc_keys (0.00s)
PASS
ok      github.com/coreybrandon/timestamp-microservice  0.234s
```

### Run

```
./timestamp_microservice
```
