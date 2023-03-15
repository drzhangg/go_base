package main

import (
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
)

func main() {
	tracer.Start(
		tracer.WithService("test-datadog1"),
		tracer.WithEnv("env"),
	)
	defer tracer.Stop()

	// Create a traced mux router
	mux := httptrace.NewServeMux()
	// Continue using the router as you normally would.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8080", mux)
}
