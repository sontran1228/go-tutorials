package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Son Tran")
	want := "Hello, Son Tran"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
