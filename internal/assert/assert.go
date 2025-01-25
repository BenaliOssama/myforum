package assert

import (
	"testing"
)

// generic function
/*To support values of either type, that single function will need a way to declare
what types it supports. Calling code, on the other hand,
will need a way to specify whether it is calling with an integer or float map.*/
// link: https://go.dev/doc/tutorial/generics
func Equal[T comparable](t *testing.T, actual, expected T) {
	// 	The t.Helper() function that we’re using in the code above indicates to the Go
	// test runner that our Equal() function is a test helper. This means that when t.Errorf()
	// is called from our Equal() function, the Go test runner will report the filename and line
	// number of the code which called our Equal() function in the output.
	t.Helper()
	if actual != expected {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}
