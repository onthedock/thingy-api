package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_newThingy(t *testing.T) {
	r := setupRouter()

	// Based on https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
	testCases := map[string]struct {
		uri        string
		wantedCode int
	}{
		"happy path":        {uri: fmt.Sprintf("/api/v1/thingy/name/%s", "newThingy"), wantedCode: http.StatusAccepted},
		"empty thingy name": {uri: fmt.Sprintf("/api/v1/thingy/name/%s", ""), wantedCode: http.StatusNotFound},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, err := http.NewRequest("POST", tc.uri, nil)
			if err != nil {
				t.Fatalf("got error %v", err)
				return
			}
			r.ServeHTTP(w, req)

			if w.Code != tc.wantedCode {
				t.Fatalf("wanted status code %d but got %d", tc.wantedCode, w.Code)
				return
			}
		})
	}
}
