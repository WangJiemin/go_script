package main

import (
	"math/big"
)

type Matrix struct {
	rows, columns int        // the number of rows and columns.
	data          []*big.Int // the contents of the matrix as one long slice.
}

// Set lets you define the value of a matrix at the given row and
// column.

func (A *Matrix) Set(r int, c int, val *big.Int) {
	A.data[findIndex(r, c, A)] = val
}

// Get retrieves the contents of the matrix at the row and column.
func (A *Matrix) Get(r, c int) *big.Int {
	return A.data[findIndex(r, c, A)]
}

// Column returns a slice that represents a column from the matrix.
// This works by examining each row, and adding the nth element of
// each to the column slice.

func (A *Matrix) Column(n int) []*big.Int {
	col := make([]*big.Int, A.rows)
	for i := 1; i <= A.rows; i++ {
		col[i-1] = A.Row(i)[n-1]
	}
	return col
}

// Row returns a slice that represents a row from the matrix.
func (A *Matrix) Row(n int) []*big.Int {
	return A.data[findIndex(n, 1, A):findIndex(n, A.columns+1, A)]
}

// Multiply multiplies two matrices together and return the resulting matrix.
// For each element of the result matrix, we get the dot product of the
// corresponding row from matrix A and column from matrix B.
func Multiply(A, B Matrix) *Matrix {
	C := Zeros(A.rows, B.columns)
	for r := 1; r <= C.rows; r++ {
		A_row := A.Row(r)
		for c := 1; c <= C.columns; c++ {
			B_col := B.Column(c)
			C.Set(r, c, dotProduct(A_row, B_col))

		}
	}
	return &C
}

// Identity creates an identity matrix with n rows and n columns.  When you
// multiply any matrix by its corresponding identity matrix, you get the
// original matrix.  The identity matrix looks like a zero-filled matrix with
// a diagonal line of one's starting at the upper left.
func Identity(n int) Matrix {
	A := Zeros(n, n)
	for i := 0; i < len(A.data); i += (n + 1) {
		A.data[i] = big.NewInt(1)
	}
	return A
}

// Zeros creates an r x c sized matrix that's filled with zeros.  The initial
// state of an int is 0, so we don't have to do any initialization.
func Zeros(r, c int) Matrix {
	return Matrix{r, c, make([]*big.Int, r*c)}
}

// New creates an r x c sized matrix that is filled with the provided data.
// The matrix data is represented as one long slice.

func New(r, c int, data []*big.Int) Matrix {
	if len(data) != r*c {
		panic("[]*big.Int data provided to matrix.New is great than the provided capacity of the matrix!'")
	}
	A := Zeros(r, c)
	A.data = data
	return A
}

// findIndex takes a row and column and returns the corresponding index
// from the underlying data slice.

func findIndex(r, c int, A *Matrix) int {
	return (r-1)*A.columns + (c - 1)
}

// dotProduct calculates the algebraic dot product of two slices.  This is just
// the sum  of the products of corresponding elements in the slices.  We use
// this when we multiply matrices together.

func dotProduct(a, b []*big.Int) *big.Int {
	total := new(big.Int)
	x := new(big.Int)
	z := new(big.Int)

	for i := 0; i < len(a); i++ {

		y := x.Mul(a[i], b[i])
		total = z.Add(total, y)
		//	total = total.Add(total, total.Mul(a[i], b[i]))
	}
	return total
}
