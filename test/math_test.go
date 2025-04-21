package math

import (
	"testing"

	"github.com/cxncxl/gogame/internal/math"
)

func TestMatrixFromSlice(t *testing.T) {
    vals := [][]float64{
        { 1, 2, 3 },
        { 1, 2 },
    }

    _, err := math.MatrixFromSlice(vals)
    if err == nil {
        t.Errorf("Expected to get an error when creating a matrix from a slice with inconsistent col length")
    }
}

func TestMatricesAdd(t *testing.T) {
    m1vs := [][]float64{
        { 1, 2, 3 },
        { 4, 5, 6 },
    }

    m2vs := [][]float64{
        { 2, 2, 2 },
        { 5, 5, 5 },
    }

    m1, err := math.MatrixFromSlice(m1vs)
    if err != nil {
        t.Error("Failed to create a matrix from a slice")
    }
    m2, err := math.MatrixFromSlice(m2vs)
    if err != nil {
        t.Error("Failed to create a matrix from a slice")
    }

    mrows := len(m1vs)
    mcols := len(m1vs[0])

    rows, cols := m1.Size()

    if mrows != rows || mcols != cols {
        t.Errorf(
            "Matrix created with a wrong size. Expected: %dx%d, got: %dx%d",
            mrows,
            mcols,
            rows,
            cols,
        )
    }

    m1.Add(m2)

    for i := range rows {
        for j := range cols {
            if m1.Get(i, j) != m1vs[i][j] + m2vs[i][j] {
                t.Errorf(
                    "Adding matrices failed. Expected: %f, got: %f", 
                    m1.Get(i, j),
                    m1vs[i][j] + m2vs[i][j],
                )
            }
        }
    }
}

func TestMatrixAddScalar(t *testing.T) {
    vals := [][]float64{
        { 1, 2, 3, 4 },
        { 5, 6, 7, 8 },
        { 9, 0, 1, 2 },
    }

    m, _ := math.MatrixFromSlice(vals)

    f := 69.0

    m.AddScalar(f)

    rows, cols := m.Size()

    for i := range rows {
        for j := range cols {
            if m.Get(i, j) != vals[i][j] + f {
                t.Errorf(
                    "Matrix + float is wrong. Expected: %f, got: %f",
                    vals[i][j] + f,
                    m.Get(i, j),
                )
            }
        }
    }
}

func TestMatrixMul(t *testing.T) {
    vals1 := [][]float64{
        { 1, 2, 3 },
        { 4, 5, 6 },
    }

    vals2 := [][]float64{
        { 7, 8 },
        { 9, 10 },
        { 11, 12 },
    }

    m1, _ := math.MatrixFromSlice(vals1)
    m2, _ := math.MatrixFromSlice(vals2)

    res, _ := m1.Mul(m2)

    correct, _ := math.MatrixFromSlice(
        [][]float64{
            { 58, 64 },
            { 139, 154 },
        },
    )

    if res.Rows() != correct.Rows() || res.Cols() != correct.Cols() {
        t.Errorf(
            "Mul returned matrix of side %dx%d, expected %dx%d",
            res.Rows(),
            res.Cols(),
            correct.Rows(),
            correct.Cols(),
        )
    }

    if !res.Equal(correct) {
        t.Error("Matrices are not equal")
    }
}

func TestMatrixCol(t *testing.T) {
    vals := [][]float64{
        { 1, 2, 3, 4 },
        { 2, 3, 4, 5 },
        { 3, 4, 5, 6 },
    }

    m, _ := math.MatrixFromSlice(vals)

    c := m.Col(1)

    if c[0] != 2 || c[1] != 3 || c[2] != 4 {
        t.Errorf("Invalid matrix col getter; got: %v", c)
    }
}
