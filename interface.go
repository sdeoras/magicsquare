package magicsquare

// MagicSquarer is an interface that generates a magic square
type MagicSquarer interface {
	// MagicSquare generates magic square
	MagicSquare(n int) ([][]int, error)
}

func NewMagicSquarer() MagicSquarer {
	return newMagicSquare()
}

func New(n int) ([][]int, error) {
	m := NewMagicSquarer()
	return m.MagicSquare(n)
}
