package fetch

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBody(t *testing.T) {
	want := "content"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(want))
	}))
	defer server.Close()

	result := <-Body(server.URL, Options{})

	got := result.Value

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
