package assertions

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type Executable func() (string, bool)

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

func AssertBoolEquals(tb testing.TB, expected, actual bool, msg string) {
	if expected != actual {
		tb.Fatalf(msg+" - expected=%t, actual=%t", expected, actual)
	}
}

func AssertStringEquals(tb testing.TB, expected, actual, msg string) {
	if expected != actual {
		tb.Fatalf(msg+" - expected=%s, actual=%s", expected, actual)
	}
}

func AssertStructEquals(tb testing.TB, expected, actual interface{}, msg string) {
	if !reflect.DeepEqual(expected, actual) {
		tb.Fatalf(msg+" - expected=%+v, actual=%+v", expected, actual)
	}
}

func AssertAll(tb testing.TB, msg string, executables ...Executable) {
	var failures []string
	for i, executable := range executables {
		if err, ok := executable(); ok {
			failures = append(failures, fmt.Sprintf("test[%d] - %s", i, err))
		}
	}

	if len(failures) != 0 {
		tb.Fatalf(fmt.Sprintf("%s (%d failures)\n", msg, len(failures)) + strings.Join(failures, "\n"))
	}
}

/*****************************************************************************
 *                             PRIVATE FUNCTIONS                             *
 *****************************************************************************/
