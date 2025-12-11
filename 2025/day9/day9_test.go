package main

import "testing"

func TestContains(t *testing.T) {
	type point struct {
		x, y int
	}

	// Define a square polygon
	square := Polygon{
		{x: 0, y: 0},
		{x: 0, y: 10},
		{x: 10, y: 10},
		{x: 10, y: 0},
	}

	tests := []struct {
		name     string
		polygon  Polygon
		point    point
		expected bool
	}{
		// Points inside
		{"Inside center", square, point{5, 5}, true},
		{"Inside near edge", square, point{1, 1}, true},

		// Points on vertices
		{"Vertex bottom-left", square, point{0, 0}, true},
		{"Vertex top-right", square, point{10, 10}, true},

		// Points on edges
		{"Edge bottom", square, point{5, 0}, true},
		{"Edge left", square, point{0, 5}, true},
		{"Edge top", square, point{5, 10}, true},
		{"Edge right", square, point{10, 5}, true},

		// Points outside
		{"Outside left", square, point{-1, 5}, false},
		{"Outside right", square, point{11, 5}, false},
		{"Outside top", square, point{5, 11}, false},
		{"Outside bottom", square, point{5, -1}, false},
		{"Outside corner", square, point{-1, -1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.polygon.Contains(tt.point.y, tt.point.x)
			if got != tt.expected {
				t.Errorf("Contains(%d, %d) = %v; want %v", tt.point.y, tt.point.x, got, tt.expected)
			}
		})
	}
}
