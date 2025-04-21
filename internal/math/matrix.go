package math

import (
	"strconv"
	"strings"
)

type Matrix struct {
    values [][]float64
}

func NewMatrix(rows int, cols int) Matrix {
    mrows := make([][]float64, rows)

    for i := range rows {
        col := make([]float64, cols)
        mrows[i] = col
    }

    return Matrix{
        values: mrows,
    }
}

func NewIdentMatrix(size int) Matrix {
    m := NewMatrix(size, size)

    for i := range size {
        m.Set(i, i, 1)
    }

    return m
}

func MatrixFromSlice(slice [][]float64) (Matrix, error) {
    prevLen := len(slice[0])
    for i := range len(slice) {
        if len(slice[i]) != prevLen {
            return Matrix{}, NotConsistentColSizeError
        }
    }

    m := NewMatrix(len(slice), len(slice[0]))

    for i := range len(slice) {
        for j := range len(slice[i]) {
            m.Set(i, j, slice[i][j])
        }
    }

    return m, nil
}

func (self Matrix) Set(row int, col int, val float64) {
    self.values[row][col] = val
}

func (self Matrix) Copy() Matrix {
    rows, cols := self.Size()

    mrows := make([][]float64, rows)

    for i := range rows {
        col := make([]float64, cols)

        for j := range cols {
            col[j] = self.values[i][j]
        }

        mrows[i] = col
    }

    return Matrix{
        values: mrows,
    }
}

func (self Matrix) Get(row int, col int) float64 {
    return self.values[row][col]
}

func (self Matrix) Value() [][]float64 {
    return self.values
}

func (self Matrix) Rows() int {
    return len(self.values)
}

func (self Matrix) Cols() int {
    return len(self.values[0])
}

func (self Matrix) Col(idx int) []float64 {
    res := make([]float64, self.Rows())

    for i := range self.Rows() {
        res[i] = self.values[i][idx]
    }

    return res
}

func (self Matrix) Row(i int) []float64 {
    return self.values[i]
}

func (self Matrix) Size() (int, int) {
    return self.Rows(), self.Cols()
}

func (self Matrix) Equal(other Matrix) bool {
    if self.Rows() != other.Rows() {
        return false
    }

    if self.Cols() != other.Cols() {
        return false
    }

    for i := range self.Rows() {
        for j := range self.Cols() {
            if self.Get(i, j) != other.Get(i, j) {
                return false
            }
        }
    }

    return true
}

func (self Matrix) Add(other Matrix) error {
    rows, cols := self.Size()
    oRows, oCols := other.Size()
    if rows != oRows || cols != oCols {
        return MatriciesOfDifferentSizeError
    }

    for i := range rows {
        for j := range cols {
            self.values[i][j] += other.Get(i, j)
        }
    }

    return nil
}

func (self Matrix) AddScalar(other float64) {
    rows, cols := self.Size()

    for i := range rows {
        for j := range cols {
            self.Set(i, j, self.Get(i, j) + other)
        }
    }
}

func (self Matrix) Sub(other Matrix) error {
    rows, cols := self.Size()
    oRows, oCols := other.Size()
    if rows != oRows || cols != oCols {
        return MatriciesOfDifferentSizeError
    }

    for i := range rows {
        for j := range cols {
            self.values[i][j] -= other.Get(i, j)
        }
    }

    return nil
}

func (self Matrix) SubScalar(other float64) {
    rows, cols := self.Size()

    for i := range rows {
        for j := range cols {
            self.Set(i, j, self.Get(i, j) - other)
        }
    }
}

func (self Matrix) Mul(other Matrix) (Matrix, error) {
    if self.Cols() != other.Rows() {
        return Matrix{}, InvalidMatricesSizeForMultiplicationError
    }

    res := NewMatrix(self.Rows(), other.Cols())

    for i := range self.Rows() {
        for j := range other.Cols() {
            col := other.Col(j)

            dot, _ := Dot(self.Row(i), col)

            res.Set(i, j, dot)
        }
    }

    return res, nil
}

func (self Matrix) MulScalar(other float64) {
    rows, cols := self.Size()

    for i := range rows {
        for j := range cols {
            self.Set(i, j, self.Get(i, j) * other)
        }
    }
}

func (self Matrix) DivScalar(other float64) {
    rows, cols := self.Size()

    for i := range rows {
        for j := range cols {
            self.Set(i, j, self.Get(i, j) / other)
        }
    }
}

func (self Matrix) String() string {
    res := "["

    rows, cols := self.Size()

    for i := range rows {
        formattedRow := make([]string, cols)

        for j := range cols {
            formattedRow[j] = strconv.FormatFloat(
                self.values[i][j],
                'f',
                -1,
                64,
            )
        }

        if i > 0 {
            res += " "
        }
        res += strings.Join(formattedRow, ", ")
        if i < rows - 1 {
            res += "\n"
        }
    }

    res += "]"

    return res
}

type notConsistentColSize struct {}
func (notConsistentColSize) Error() string {
    return "Not consistent column size! All columns in matrix must be of the same size"
}
var NotConsistentColSizeError = notConsistentColSize{}

type matriciesOfDifferentSize struct {}
func (matriciesOfDifferentSize) Error() string {
    return "Matricies have different size"
}
var MatriciesOfDifferentSizeError = matriciesOfDifferentSize{}

type invalidMatricesSizeForMultiplication struct {}
func (invalidMatricesSizeForMultiplication) Error() string {
    return "Number of columns on self must equal to number of rows of other"
}
var InvalidMatricesSizeForMultiplicationError = invalidMatricesSizeForMultiplication{}
