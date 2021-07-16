package main

import "testing"
import "bytes"

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nhan")

	got := buffer.String()
	want := "Hello, Nhan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
