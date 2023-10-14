package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	var target, got, want string

	target = "target"
	want = "foo: target"
	got = foo(target)
	if got != want {
		t.Errorf("foo(%v) == %v, but wants %v", target, got, want)
	}
}
