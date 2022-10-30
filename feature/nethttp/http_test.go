package nethttp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	http.HandleFunc("/", handler)
	t.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "request: %s\n", r.URL.Path)
	fmt.Fprintf(w, "append1: %s\n", r.URL.Path)
	fmt.Fprintf(w, "append2: %s\n", r.URL.Path)
}
