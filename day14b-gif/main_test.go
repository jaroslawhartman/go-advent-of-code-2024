package main

import "testing"

func TestRun(t *testing.T) {
	got := run("input_test.txt", 11, 7)
	expected := 12
	if got != expected {
		t.Errorf("Run = %d; want %d", got, expected)
	}
}
