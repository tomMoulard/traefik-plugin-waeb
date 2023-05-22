package traefik_plugin_waeb_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	waeb "github.com/tomMoulard/traefik-plugin-waeb"
)

func TestDemo(t *testing.T) {
	t.Parallel()

	cfg := waeb.CreateConfig()

	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Should NEVER go through the next handler
		t.Fail()
	})

	handler, err := waeb.New(context.Background(), next, cfg, "waeb")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://localhost/waeb.go", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("invalid recorder status code, expected: %d, got: %d", http.StatusOK, recorder.Code)
	}

	if recorder.Header().Get("Content-Type") != "text/x-go; charset=utf-8" {
		t.Errorf("invalid Content-Type, expected: %q, got: %q",
			"text/x-go; charset=utf-8", recorder.Header().Get("Content-Type"))
	}
}
