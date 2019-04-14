package magicsquare

import (
	"testing"
)

func TestNew_OddOrder(t *testing.T) {
	n := 5

	y := [][]int{
		{17, 24, 1, 8, 15},
		{23, 5, 7, 14, 16},
		{4, 6, 13, 20, 22},
		{10, 12, 19, 21, 3},
		{11, 18, 25, 2, 9},
	}

	x, err := New(n)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if x[i][j] != y[i][j] {
				t.Fatal("expected:", y[i][j], "got:", x[i][j])
			}
		}
	}
}

func TestNew_DoublyEven(t *testing.T) {
	n := 4

	x, err := New(n)
	if err != nil {
		t.Fatal(err)
	}

	y := [][]int{
		{16, 2, 3, 13},
		{5, 11, 10, 8},
		{9, 7, 6, 12},
		{4, 14, 15, 1},
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if x[i][j] != y[i][j] {
				t.Fatal("expected:", y[i][j], "got:", x[i][j])
			}
		}
	}
}

func TestNew_SinglyEven(t *testing.T) {
	n := 6

	x, err := New(n)
	if err != nil {
		t.Fatal(err)
	}

	y := [][]int{
		{35, 1, 6, 26, 19, 24},
		{3, 32, 7, 21, 23, 25},
		{31, 9, 2, 22, 27, 20},
		{8, 28, 33, 17, 10, 15},
		{30, 5, 34, 12, 14, 16},
		{4, 36, 29, 13, 18, 11},
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if x[i][j] != y[i][j] {
				t.Fatal("expected:", y[i][j], "got:", x[i][j])
			}
		}
	}
}
