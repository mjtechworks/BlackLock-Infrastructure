

import (
    "net/http"
    "context"
    "testing"
    "net/http/httptest"
    "bufio"
    "bytes"
    "fmt"
    "net/http"
    "os"
    "strings"
)

func NewContextWithRequestID(ctx context.Context, r *http.Request) context.Context {
    return context.WithValue(ctx, "reqId", "1234")
}

func AddContextWithRequestID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var ctx = context.Background()
        ctx = NewContextWithRequestID(ctx, r)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func TestIt(t *testing.T) {

    // create a handler to use as "next" which will verify the request
    nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        val := r.Context().Value("reqId")
        if val == nil {
            t.Error("reqId not present")
        }
        valStr, ok := val.(string)
        if !ok {
            t.Error("not string")
        }
        if valStr != "1234" {
            t.Error("wrong reqId")
        }
    })

    // create the handler to test, using our custom "next" handler
    handlerToTest := AddContextWithRequestID(nextHandler)

    // create a mock request to use
    req := httptest.NewRequest("GET", "http://testing", nil)

    // call the handler using a mock response recorder (we'll not use that anyway)
    handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
}

func Timer(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        timer := time.Now()
        next.ServeHTTP(w, r)

        if id, ok := r.Context().Value("some-key").(string); ok {
            // Do something with the id and timer
        }
    }
}
