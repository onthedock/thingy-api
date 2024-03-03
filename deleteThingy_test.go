package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oklog/ulid/v2"
)

func Test_deleteThingy(t *testing.T) {
	thingyID := ulid.Make()
	testCases := map[string]struct {
		uri      string
		wantCode int
	}{
		"happy path":            {uri: fmt.Sprintf("/api/v1/thingy/id/%s", thingyID.String()), wantCode: http.StatusAccepted},
		"empty thingyId":        {uri: fmt.Sprintf("/api/v1/thingy/id/%s", ""), wantCode: http.StatusNotFound},
		"invalid thingyId":      {uri: fmt.Sprintf("/api/v1/thingy/id/%s", "123"), wantCode: http.StatusBadRequest},
		"non-existing thingyId": {uri: fmt.Sprintf("/api/v1/thingy/id/%s", ulid.Make().String()), wantCode: http.StatusGone},
	}

	r := setupRouter()
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if _, err := addThingyToDB(&Thingy{Id: thingyID, Name: "testThingy"}); err != nil {
				t.Fatalf("failed to add test Thingy to DB: %v", err)
				return
			}
			resp := httptest.NewRecorder()
			req, err := http.NewRequest("DELETE", tc.uri, nil)
			if err != nil {
				t.Fatalf("got error %v", err)
				return
			}
			r.ServeHTTP(resp, req)
			if resp.Code != tc.wantCode {
				t.Fatalf("wanted %d but got %d: %v", tc.wantCode, resp.Code, resp.Body)
				return
			}
		})
	}

}
