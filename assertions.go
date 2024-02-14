package assertions

import (
	"fmt"
	"reflect"
	"testing"
)

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

func AssertNotNull(tb testing.TB, val any, msg string) {
	if isNil(val) {
		tb.Fatal(msg + fmt.Sprintf(" - expected not %+v", val))
	}
}

func AssertBoolEquals(tb testing.TB, expected, actual bool, msg string) {
	if expected != actual {
		tb.Fatal(msg + fmt.Sprintf(" - expected=%t, actual=%t", expected, actual))
	}
}

func AssertIntEquals(tb testing.TB, expected, actual int, msg string) {
	if expected != actual {
		tb.Fatal(msg + fmt.Sprintf(" - expected=%d, actual=%d", expected, actual))
	}
}

func AssertInt64Equals(tb testing.TB, expected, actual int64, msg string) {
	if expected != actual {
		tb.Fatal(msg + fmt.Sprintf(" - expected=%d, actual=%d", expected, actual))
	}
}

func AssertStringEquals(tb testing.TB, expected, actual, msg string) {
	if expected != actual {
		tb.Fatal(msg + fmt.Sprintf(" - expected=%s, actual=%s", expected, actual))
	}
}

func AssertEquals(tb testing.TB, expected, actual any, msg string) {
	if expected != actual {
		tb.Fatal(msg + fmt.Sprintf(" - expected=%+v, actual=%+v", expected, actual))
	}
}

func AssertDeepEquals(tb testing.TB, expected, actual any, msg string) {
	if !reflect.DeepEqual(expected, actual) {
		tb.Fatal(msg + fmt.Sprintf(" - expected=%+v, actual=%+v", expected, actual))
	}
}

/*****************************************************************************
 *                             PRIVATE FUNCTIONS                             *
 *****************************************************************************/

func isNil(val any) bool {
	ret := false

	if val == nil {
		ret = true
	}

	switch reflect.TypeOf(val).Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		ret = reflect.ValueOf(val).IsNil()
	}
	return ret
}
