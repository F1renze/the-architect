package test

import (
	"reflect"
	"testing"
)

func Assert(tb testing.TB, condition bool, format string, v ...interface{}) {
	tb.Helper()
	if condition != true {
		tb.Fatalf(format, v...)
	}
}

func AssertEqual(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("excepted '%v', got '%v'", exp, act)
	}
}

func NoErr(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("unexpected error occur, '%v'", err)
	}
}
