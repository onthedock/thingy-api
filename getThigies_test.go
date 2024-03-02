package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getThingies(t *testing.T) {
	r := setupRouter()
	type testCase struct {
		method      string
		uri         string
		wantedCode  int
		wantedError error
	}

	testCases := []testCase{
		{method: "GET", uri: "/api/v1/thingy", wantedCode: http.StatusOK, wantedError: nil},
	}

	w := httptest.NewRecorder()

	for _, tc := range testCases {
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
	}
}
