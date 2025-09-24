package main

import "testing"

func TestReverseStr(t *testing.T) {
	got := ReverseStr("Hello goland")
	want := "dnalog olleH"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
