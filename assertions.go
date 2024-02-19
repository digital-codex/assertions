package assertions

import (
	"reflect"
	"testing"
)

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

func AssertNotNull(tb testing.TB, actual any, msg string) {
	if isNil(actual) {
		tb.Fatalf("%s: expected not <%v>", msg, actual)
	}
}

func AssertNotEquals(tb testing.TB, unexpected, actual any, msg string) {
	if unexpected == actual {
		tb.Fatalf("%s: not equal but was <%v>", msg, unexpected)
	}
}

func AssertBoolEquals(tb testing.TB, expected, actual bool, msg string) {
	if expected != actual {
		tb.Fatalf("%s: expected=%t, actual=%t", msg, expected, actual)
	}
}

func AssertIntEquals(tb testing.TB, expected, actual int, msg string) {
	if expected != actual {
		tb.Fatalf("%s: expected=%d, actual=%d", msg, expected, actual)
	}
}

func AssertInt64Equals(tb testing.TB, expected, actual int64, msg string) {
	if expected != actual {
		tb.Fatalf("%s: expected=%d, actual=%d", msg, expected, actual)
	}
}

func AssertFloat64Equals(tb testing.TB, expected, actual float64, msg string) {
	if expected != actual {
		tb.Fatalf("%s: expected=%d, actual=%d", msg, expected, actual)
	}
}

func AssertStringEquals(tb testing.TB, expected, actual, msg string) {
	if expected != actual {
		tb.Fatalf("%s: expected=%q, actual=%q", msg, expected, actual)
	}
}

func AssertEquals(tb testing.TB, expected, actual any, msg string) {
	if expected != actual {
		tb.Fatalf("%s: expected=%v, actual=%v", msg, expected, actual)
	}
}

func AssertDeepEquals(tb testing.TB, expected, actual any, msg string) {
	if !reflect.DeepEqual(expected, actual) {
		tb.Fatalf("%s: expected=%v, actual=%v", msg, expected, actual)
	}
}

func AssertTypeOf(tb testing.TB, expected reflect.Type, actual any, msg string) {
	rt := reflect.TypeOf(actual)
	if rt.Kind() == reflect.Interface || rt.Kind() == reflect.Ptr {
		rt = reflect.ValueOf(actual).Elem().Type()
	}

	switch expected.Kind() {
	case reflect.Struct:
		if expected != rt {
			tb.Fatalf("%s: expected=%s, actual=%s", msg, expected.Name(), rt.Name())
		}
		// NOTE: I'm not sure if the other checks are necessary here
		if expected.NumField() != rt.NumField() {
			tb.Fatalf("%s: NumField() wrong: expected=%d, actual=%d", msg, expected.NumField(), rt.NumField())
		}
		for i := 0; i < expected.NumField(); i += 1 {
			if !expected.Field(i).IsExported() && expected.Field(i).PkgPath != rt.Field(i).PkgPath {
				tb.Fatalf("%s: Field(%d).PkgPath on non-exported field wrong: expected=%s, actual=%s", msg, i, expected.Field(i).PkgPath, rt.Field(i).PkgPath)
			}
			if expected.Field(i).Name != rt.Field(i).Name {
				tb.Fatalf("%s: Field(%d).Name wrong: expected=%s, actual=%s", msg, i, expected.Field(i).Name, rt.Field(i).Name)
			}
			if expected.Field(i).Type != rt.Field(i).Type {
				tb.Fatalf("%s: Field(%d).Type wrong: expected=%s, actual=%s", msg, i, expected.Field(i).Type.Name(), rt.Field(i).Type.Name())
			}
			if expected.Field(i).Tag != rt.Field(i).Tag {
				tb.Fatalf("%s: Field(%d).Tag wrong: expected=%s, actual=%s", msg, i, expected.Field(i).Tag, rt.Field(i).Tag)
			}
		}
	case reflect.Ptr:
		if expected.Elem() != rt {
			tb.Fatalf("%s: expected=%s, actual=%s", msg, expected.Elem().Name(), rt.Name())
		}
	case reflect.Interface:
		if !rt.Implements(expected) {
			tb.Fatalf("%s: expected=%s, actual=%s", msg, expected.Name(), rt.Name())
		}
	default:
		panic("unsupported operation")
	}
}

/*****************************************************************************
 *                             PRIVATE FUNCTIONS                             *
 *****************************************************************************/

func isNil(i any) bool {
	ret := false

	if i == nil {
		ret = true
	}

	if i != nil {
		switch reflect.TypeOf(i).Kind() {
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
			ret = reflect.ValueOf(i).IsNil()
		default:
			ret = false
		}
	}

	return ret
}
