package main

import "fmt"
import "testing"

func TestHelloWorld(t *testing.T) {
	cases := []struct {
		A, B, Expected int
	}{
		{1, 1, 2},
		{1, -1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d + %d", tc.A, tc.B), func(t *testing.T) {
			actual := tc.A + tc.B
			if actual != tc.Expected {
				t.Fatal("failure")
			}
		})
	}
}

func TestFib(t *testing.T) {
	cases := []struct {
		n, expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("fib(%d)", tc.n), func(t *testing.T) {
			actual := fib(tc.n)
			if actual != tc.expected {
				t.Fatalf("fib(%d):  Expected: %d, Was: %d", tc.n, tc.expected, actual)
			}
		})
	}
}

func TestFibList(t *testing.T) {
	cases := []struct {
		n        int
		expected []int
	}{
		{1, []int{1}},
		{2, []int{1, 1}},
		{3, []int{1, 1, 2}},
		{4, []int{1, 1, 2, 3}},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("fib(%d)", tc.n), func(t *testing.T) {
			actual := fiblist(tc.n)
			if len(actual) != len(tc.expected) {
				t.Fatal("Expected: %v, Was: %v", tc.expected, actual)
			}
			for index := range actual {
				if actual[index] != tc.expected[index] {
					t.Fatalf("Expected: %v, Was: %v", tc.expected, actual)
				}
			}
		})
	}
}
