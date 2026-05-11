package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// result := Add(2, 3)
	// fmt.Printf("Test 1 result is %d\n", result)

	tests := []struct {
		a, b, expected int64
	}{
		{1, 2, 3},
		{3, 2, 5},
		{11, 20, 30},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
