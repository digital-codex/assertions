package assertions

import (
	"fmt"
	"reflect"
	"testing"
)

type Executable func() (string, bool)

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

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

func AssertStructEquals(tb testing.TB, expected, actual *interface{}, msg string) {
	if expected != actual {
		tb.Fatal(msg + fmt.Sprintf(" - expected=%+v, actual=%+v", expected, actual))
	}
}

func AssertDeepEquals(tb testing.TB, expected, actual interface{}, msg string) {
	if !reflect.DeepEqual(expected, actual) {
		tb.Fatal(msg + fmt.Sprintf(" - expected=%+v, actual=%+v", expected, actual))
	}
}

/*****************************************************************************
 *                             PRIVATE FUNCTIONS                             *
 *****************************************************************************/
