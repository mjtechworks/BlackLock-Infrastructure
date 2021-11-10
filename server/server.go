package main

import (
    "fmt"
    "log"                                                                                                    
    "net/http"
)

type StatusRecorder struct {
    http.ResponseWriter
    Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
    r.Status = status
    r.ResponseWriter.WriteHeader(status)
}

func WithLogging(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        recorder := &StatusRecorder{
            ResponseWriter: w,
            Status:         200,
        }
        h.ServeHTTP(recorder, r)
        log.Printf("Handling request for %s from %s, status: %d", r.URL.Path, r.RemoteAddr, recorder.Status)
    })
}

func main() {
    // API routes
    myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %s", r.RemoteAddr)
    })
    handlerWithLogging := WithLogging(myHandler)
    http.Handle("/", handlerWithLogging)
    http.ListenAndServe(":2000", nil)
}
    
func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path == "/foo" {
			return
		}

		next.ServeHTTP(w, r)
		log.Println("Executing middlewareTwo again")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", middlewareOne(middlewareTwo(finalHandler)))

	log.Println("Listening on :2000...")
	err := http.ListenAndServe(":2000", mux)
	log.Fatal(err)
}
