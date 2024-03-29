package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getThingies(t *testing.T) {
	r := setupRouter()

	// Based on https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
	testCases := map[string]struct {
		method      string
		uri         string
		wantedCode  int
		wantedError error
	}{
		"happy path":      {method: "GET", uri: "/api/v1/thingy", wantedCode: http.StatusOK, wantedError: nil},
		"negative offset": {method: "GET", uri: "/api/v1/thingy?offset=-3", wantedCode: http.StatusOK, wantedError: nil},
		"offset bigger that length of items in db":             {method: "GET", uri: "/api/v1/thingy?offset=3000", wantedCode: http.StatusOK, wantedError: nil},
		"offset + thigiesPerPage exceed length of items in db": {method: "GET", uri: "/api/v1/thingy?offset=30", wantedCode: http.StatusOK, wantedError: nil},
		"invalid offset (not a number)":                        {method: "GET", uri: "/api/v1/thingy?offset=patata", wantedCode: http.StatusOK, wantedError: nil},
	}

	w := httptest.NewRecorder()

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, tc.uri, nil)
			if err != tc.wantedError {
				t.Fatalf("wanted error %v but got %v", tc.wantedError, err)
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
