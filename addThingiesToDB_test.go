package main

import "testing"

func Test_addThingiesToDB(t *testing.T) {
	// Based on https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
	testCases := map[string]struct {
		numberOfThingies int
		want             int
	}{
		"happy path":      {numberOfThingies: 35, want: 35},
		"negative number": {numberOfThingies: -3, want: 5},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			var localThingies []Thingy = addThingiesToDB(tc.numberOfThingies)

			if len(localThingies) != tc.want {
				t.Fatalf("wanted %d thingies but got %d", tc.want, len(localThingies))
				return
			}
		})
	}
}
