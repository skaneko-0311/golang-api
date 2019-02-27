package main

import "testing"

func TestAddSuccess(t *testing.T) {
	result, err := add(1, 2)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 3 {
		t.Fatal("failed test")
	}
}

func TestAddFailed(t *testing.T) {
	result, err := add(4, 8)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 12 {
		t.Fatal("failed test")
	}
}
