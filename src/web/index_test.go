package web

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestIndex_PrintsHelloWorld(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Fatal("Non-200 status:", w.Code, w.Body)
	}
	if w.Body.String() != "Hello, world!\n" {
		t.Fatalf("Bad body: %#v", w.Body.String())
	}
}
