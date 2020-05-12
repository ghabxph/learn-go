# Justforfunc #16: unit testing HTTP servers

## Youtube video:

- https://www.youtube.com/watch?v=hVFEV-ieeew


## Interesting parts:

- 12:00 - Elegant way of writing test table <3
- 13:00 - Moar elegant way (added message.)
- 14:26 - Subtests
- 16:10 - Examples as tests
- 20:03 - Testing http handlers
- 23:45 - net/http/httptest
- 27:41 - This is where I stopped watching
- 29:50 - How to test the routing from the server
- 39:12 - How to test the routing from the server

## Nice `go test` commands

- `go test -v` - *shows all unit test / sub test that are executed`*
- `go test -v -run <name of test>` - *executes unit test that matches the pattern*

## Nice `http` unit test snippet learned from this guy.

*main_test.go*
```go
package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
)

func TestDoubleHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "localhost:8080/doulbe?v=2", nil)
    if err != nil {
        t.Fatalf("could not create request: %v", err)
    }
    rec := httptest.NewRecorder()

    doubleHandler(rec, req)

    res := rec.Result()
    if res.StatusCode != http.StatusOK {
        t.Errorf("expected statuis OK; got %v", res.Status)
    }
}

```
