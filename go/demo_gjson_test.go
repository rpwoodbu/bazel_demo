package main

import "testing"

func TestComputeHelloBasic(t *testing.T) {
	input := `{"greet": "Bazel Rocks"}`
	expected := "Hello, Bazel Rocks!"
	actual := compute_hello(input)
	if actual != expected {
		t.Errorf("compute_hello(%s) = %q; want %q", input, actual, expected)
	}
}

func TestComputeHelloEmptyValue(t *testing.T) {
	input := `{"greet": ""}`
	expected := "Hello, !"
	actual := compute_hello(input)
	if actual != expected {
		t.Errorf("compute_hello(%s) = %q; want %q", input, actual, expected)
	}
}

func TestComputeHelloEmptyJson(t *testing.T) {
	input := ``
	expected := "Hello, !"
	actual := compute_hello(input)
	if actual != expected {
		t.Errorf("compute_hello(%s) = %q; want %q", input, actual, expected)
	}
}
