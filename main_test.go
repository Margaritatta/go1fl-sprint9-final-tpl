package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{"zero size", 0, 0},
		{"negative size", -5, 0},
		{"one element", 1, 1},
		{"many elements", 1000, 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.wantLen {
				t.Errorf("len(got) = %d, want %d", len(got), tt.wantLen)
			}
			for _, v := range got {
				if v < 0 {
					t.Errorf("got negative element %d", v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"nil slice", nil, 0},
		{"single element", []int{42}, 42},
		{"max at beginning", []int{9, 3, 5}, 9},
		{"max in middle", []int{1, 8, 3}, 8},
		{"max at end", []int{1, 3, 7}, 7},
		{"all equal", []int{5, 5, 5}, 5},
		{"large values", []int{1, 9223372036854775807, 2}, 9223372036854775807},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"nil slice", nil, 0},
		{"single element", []int{10}, 10},
		{"less than chunks", []int{1, 2, 3}, 3},
		{"exact chunks", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"max in last chunk", []int{1, 2, 3, 4, 5, 6, 7, 9}, 9},
		{"all equal", []int{4, 4, 4, 4, 4, 4, 4, 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	data := generateRandomElements(10_000)
	want := maximum(data)
	if got := maxChunks(data); got != want {
		t.Errorf("maxChunks() = %d, want %d", got, want)
	}
}
