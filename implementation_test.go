package magicsquare

import (
	"testing"
)

// TestNew_Many generates 100 magic squares and verifies them
func TestNew_Many(t *testing.T) {
	for n := 3; n < 100; n++ {
		x, err := New(n)
		if err != nil {
			t.Fatal(err)
		}

		if !verify(x) {
			t.Fatal("sum verification failed for:", n)
		}
	}
}

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

	if !verify(x) {
		t.Fatal("sum verification failed")
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

	if !verify(x) {
		t.Fatal("sum verification failed")
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

	if !verify(x) {
		t.Fatal("sum verification failed")
	}
}

func verify(x [][]int) bool {
	n := len(x)

	sum := 0
	forwardSlashDiagSum := 0
	backSlashDiagSum := 0

	for i := 0; i < n; i++ {
		forwardSlashDiagSum += x[n-i-1][i]
		backSlashDiagSum += x[i][i]

		rowSum := 0
		colSum := 0
		for j := 0; j < n; j++ {
			rowSum += x[i][j]
			colSum += x[j][i]
		}

		if sum == 0 {
			sum = rowSum
		} else {
			if rowSum != sum {
				return false
			}
		}

		if colSum != sum {
			return false
		}
	}

	if forwardSlashDiagSum != sum ||
		backSlashDiagSum != sum {
		return false
	}

	return true
}
