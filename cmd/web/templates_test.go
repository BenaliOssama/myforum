package main

import (
	"myforum/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2022 at 10:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2022 at 09:15",
		},
	}

	// Loop over the test cases.
	for _, tt := range tests {
		// Use the t.Run() function to run a sub-test for each test case. The
		// first parameter to this is the name of the test (which is used to
		// identify the sub-test in any log output) and the second parameter is
		// and anonymous function containing the actual test for each case.
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			// if hd != tt.want {
			// 	t.Errorf("got %q; want %q", hd, tt.want)
			// }

			// Use the new assert.Equal() helper to compare the expected and
			// actual values.
			assert.Equal(t, hd, tt.want)
		})

	}
}

/*It’s important to point out that you don’t need to use sub-tests in conjunction with table-
driven tests (like we have done so far in this chapter). It’s perfectly valid to execute sub-tests
by calling t.Run() consecutively in your test functions, similar to this:
func TestExample(t *testing.T) {
	t.Run("Example sub-test 1", func(t *testing.T) {
		// Do a test.
	})
	t.Run("Example sub-test 2", func(t *testing.T) {
		// Do another test.
	})
	t.Run("Example sub-test 3", func(t *testing.T) {
		// And another...
	})
}
*/
