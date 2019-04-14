package magicsquare

func newMagicSquare() *magicSquare {
	return new(magicSquare)
}

type magicSquare struct{}

func (m *magicSquare) MagicSquare(n int) ([][]int, error) {
	if n%2 != 0 {
		return oddOrderMagicSquare(n), nil
	}

	if n%4 == 0 {
		return doublyEvenMagicSquare(n), nil
	}

	return singlyEvenMagicSquare(n), nil
}

// derived from: https://rosettacode.org/wiki/Magic_squares_of_singly_even_order#Go
func singlyEvenMagicSquare(n int) [][]int {
	size := n * n
	halfN := n / 2
	subSquareSize := size / 4
	subSquare := oddOrderMagicSquare(halfN)

	quadrantFactors := [4]int{0, 2, 3, 1}
	x := make([][]int, n)

	for i := 0; i < n; i++ {
		x[i] = make([]int, n)
	}

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			quadrant := r/halfN*2 + c/halfN
			x[r][c] = subSquare[r%halfN][c%halfN]
			x[r][c] += quadrantFactors[quadrant] * subSquareSize
		}
	}

	nColsLeft := halfN / 2
	nColsRight := nColsLeft - 1

	for r := 0; r < halfN; r++ {
		for c := 0; c < n; c++ {
			if c < nColsLeft || c >= n-nColsRight ||
				(c == nColsLeft && r == nColsLeft) {
				if c == 0 && r == nColsLeft {
					continue
				}
				tmp := x[r][c]
				x[r][c] = x[r+halfN][c]
				x[r+halfN][c] = tmp
			}
		}
	}
	return x
}

// derived from: https://www.geeksforgeeks.org/magic-square-even-order/
func doublyEvenMagicSquare(n int) [][]int {
	// allocate memory
	// assign values starting from 1
	x := make([][]int, n)
	for i := range x {
		x[i] = make([]int, n)
		for j := range x[i] {
			x[i][j] = n*i + j + 1
		}
	}

	// change value of Array elements
	// at fix location as per rule
	// (n*n+1)-arr[i][j]
	// Top Left corner of Matrix
	// (order (n/4)*(n/4))
	for i := 0; i < n/4; i++ {
		for j := 0; j < n/4; j++ {
			x[i][j] = (n*n + 1) - x[i][j]
		}
	}

	// Top Right corner of Matrix
	// (order (n/4)*(n/4))
	for i := 0; i < n/4; i++ {
		for j := 3 * (n / 4); j < n; j++ {
			x[i][j] = (n*n + 1) - x[i][j]
		}
	}

	// Bottom Left corner of Matrix
	// (order (n/4)*(n/4))
	for i := 3 * n / 4; i < n; i++ {
		for j := 0; j < n/4; j++ {
			x[i][j] = (n*n + 1) - x[i][j]
		}
	}

	// Bottom Right corner of Matrix
	// (order (n/4)*(n/4))
	for i := 3 * n / 4; i < n; i++ {
		for j := 3 * n / 4; j < n; j++ {
			x[i][j] = (n*n + 1) - x[i][j]
		}
	}

	// Centre of Matrix (order (n/2)*(n/2))
	for i := n / 4; i < 3*n/4; i++ {
		for j := n / 4; j < 3*n/4; j++ {
			x[i][j] = (n*n + 1) - x[i][j]
		}
	}

	return x
}

func oddOrderMagicSquare(n int) [][]int {
	// allocate memory
	x := make([][]int, n)
	for i := range x {
		x[i] = make([]int, n)
	}

	row := 0 + 1
	col := n/2 - 1
	for i := 0; i < n*n; i++ {
		// try to go in top right cell
		row--
		col++

		// this is either
		if row < 0 {
			if col >= 0 && col < n {
				row = n - 1
			} else {
				row += 2
				col--
			}
		} else {
			if col >= 0 && col < n {
				if x[row][col] != 0 {
					row += 2
					col--
				}
			} else {
				col = 0
			}
		}

		x[row][col] = i + 1
	}

	return x
}
