package ewords

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

func Minimise(jsn string) string {
	re, _ := regexp.Compile(`[\s\n]`)
	return strings.TrimSpace(re.ReplaceAllString(jsn, ""))
}

func AssertNil(t *testing.T, actual interface{}) {

	if reflect.ValueOf(actual).Kind() == reflect.Ptr && reflect.ValueOf(actual).IsNil() {
		t.Logf(`%s
		Expected:%v
		Actual:  %v
		`, success, nil, actual)
	} else {
		t.Fatalf(`%s
	    Expected:%v
	    Actual:  %v
	    `, failed, nil, actual)
	}

}

func AssertEqualsString(t *testing.T, expected string, actual string) {
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
