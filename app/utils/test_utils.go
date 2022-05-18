package utils

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func StrPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}

func FloatPtr(f float64) *float64 {
	return &f
}

func IntPtr(f int) *int {
	return &f
}

//TODO Need to make more clever solution for minimizing, at least in which spaces inside quotes won't be removed
func Minimise(jsn string) string {
	re, _ := regexp.Compile(`[\s\n]`)
	return strings.TrimSpace(re.ReplaceAllString(jsn, ""))
}

func AssertNil(t *testing.T, actual interface{}) {
	t.Helper()
	if actual == nil ||
		((reflect.ValueOf(actual).Kind() == reflect.Ptr ||
			reflect.ValueOf(actual).Kind() == reflect.Map ||
			reflect.ValueOf(actual).Kind() == reflect.Slice ||
			reflect.ValueOf(actual).Kind() == reflect.Chan ||
			reflect.ValueOf(actual).Kind() == reflect.Func ||
			reflect.ValueOf(actual).Kind() == reflect.UnsafePointer ||
			reflect.ValueOf(actual).Kind() == reflect.Interface) && reflect.ValueOf(actual).IsNil()) {
		t.Logf(`%s Value is nill`, success)
	} else {
		t.Fatalf(`%s Expected nil value, but: %v`, failed, PrettyPrint(actual))
	}
}

func AssertNotNill(t *testing.T, actual interface{}) {
	t.Helper()
	if (reflect.ValueOf(actual).Kind() == reflect.Ptr ||
		reflect.ValueOf(actual).Kind() == reflect.Map ||
		reflect.ValueOf(actual).Kind() == reflect.Slice ||
		reflect.ValueOf(actual).Kind() == reflect.Chan ||
		reflect.ValueOf(actual).Kind() == reflect.Func ||
		reflect.ValueOf(actual).Kind() == reflect.UnsafePointer ||
		reflect.ValueOf(actual).Kind() == reflect.Interface) && !reflect.ValueOf(actual).IsNil() {
		t.Logf(`%s %v - Not nill`, success, PrettyPrint(actual))
	} else {
		t.Fatalf(`%s Exptected not nil but: %v`, failed, PrettyPrint(actual))
	}
}

func AssertEmptyString(t *testing.T, actual string) {
	t.Helper()
	AssertEqualsString(t, "", actual)
}

func AssertEqualsString(t *testing.T, expected string, actual string) {
	t.Helper()
	if expected != actual {
		t.Fatalf(`%s
		Expected:%v
		Actual:  %v
		`, failed, expected, actual)
	}
	t.Logf(`%s
		Expected:%v
		Actual:  %v
		`, success, expected, actual)
}

func AssertEqualsBool(t *testing.T, expected bool, actual bool) {
	t.Helper()
	if expected != actual {
		t.Fatalf(`%s
		Expected:%v
		Actual:  %v
		`, failed, expected, actual)
	}
	t.Logf(`%s
		Expected:%v
		Actual:  %v
		`, success, expected, actual)
}

func AssertEqualsFloat64(t *testing.T, expected float64, actual float64) {
	t.Helper()
	if expected != actual {
		t.Fatalf(`%s
		Expected:%v
		Actual:  %v
		`, failed, expected, actual)
	}
	t.Logf(`%s
		Expected:%v
		Actual:  %v
		`, success, expected, actual)
}

func AssertEqualsInt(t *testing.T, expected int, actual int) {
	t.Helper()
	if expected != actual {
		t.Fatalf(`%s
		Expected:%v
		Actual:  %v
		`, failed, expected, actual)
	}
	t.Logf(`%s
		Expected:%v
		Actual:  %v
		`, success, expected, actual)
}

func AssertEqualsTime(t *testing.T, expected time.Time, actual time.Time) {
	t.Helper()
	if expected != actual {
		t.Fatalf(`%s
		Expected:%v
		Actual:  %v
		`, failed, expected, actual)
	}
	t.Logf(`%s
		Expected:%v
		Actual:  %v
		`, success, expected, actual)
}
