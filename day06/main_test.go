package main

import "testing"

func TestRun(t *testing.T) {
	got := run("input_test.txt")
	expected := 41
	if got != expected {
		t.Errorf("Run = %d; want %d", got, expected)
	}
}
