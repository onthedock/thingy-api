package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getThingyById(t *testing.T) {
	r := setupRouter()

	// Based on https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
	testCases := map[string]struct {
		uri         string
		wantedCode  int
		wantedError error
	}{
		"happy path":       {uri: fmt.Sprintf("/api/v1/thingy/id/%s", thingiesDB[0].Id.String()), wantedCode: http.StatusOK, wantedError: nil},
		"invalid thingy":   {uri: "/api/v1/thingy/id/123", wantedCode: http.StatusBadRequest, wantedError: nil},
		"thingy not found": {uri: "/api/v1/thingy/id/01HR0361PRQFAHBEVK2AT43GYQ", wantedCode: http.StatusNotFound, wantedError: nil},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", tc.uri, nil)
			if err != tc.wantedError {
				t.Fatalf("wanted error %v but got %v", tc.wantedError, err)
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
