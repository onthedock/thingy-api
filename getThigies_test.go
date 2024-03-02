package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getThingies(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/thingy", nil)
	if err != nil {
		t.Fatalf("unable to create request: %s", err.Error())
		return
	}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("wanted status code %d but got %d", http.StatusOK, w.Code)
		return
	}
}
