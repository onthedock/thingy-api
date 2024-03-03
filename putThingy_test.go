package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_putThingy(t *testing.T) {
	r := setupRouter()

	// Based on https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
	testCases := map[string]struct {
		uri        string
		wantedCode int
		data       io.Reader
	}{
		"happy path":      {uri: "/api/v1/thingy", wantedCode: http.StatusAccepted, data: bytes.NewBuffer([]byte(`{"id": "01HR06TEAB5HRKRG1G8J0HJ6KG", "name": "addedThingy"}`))},
		"incomplete data": {uri: "/api/v1/thingy", wantedCode: http.StatusBadRequest, data: bytes.NewBuffer([]byte(`{"id": "01HR06TEAB5HRKRG1G8J0HJ6KG"}`))},
		"invalid format":  {uri: "/api/v1/thingy", wantedCode: http.StatusBadRequest, data: bytes.NewBuffer([]byte(`{"id": "123", "name": "invalidTHingy"}`))},
		"no data":         {uri: "/api/v1/thingy", wantedCode: http.StatusBadRequest, data: nil},
	}

	for name, tc := range testCases {
		w := httptest.NewRecorder()
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("PUT", tc.uri, tc.data)
			if err != nil {
				t.Fatalf("wanted error %v but got", err)
				return
			}
			r.ServeHTTP(w, req)

			if w.Code != tc.wantedCode {
				t.Fatalf("wanted status code %d but got %d: %v", tc.wantedCode, w.Code, w.Body)
				return
			}
		})
	}
}
