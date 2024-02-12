package assertions

import "testing"

func AssertEquals(tb testing.TB, expected, actual string) {
	if expected != actual {
		tb.Fatalf("expected=%s, actual=%s", expected, actual)
	}
}
