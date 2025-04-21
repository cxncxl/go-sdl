package math

func Dot(a []float64, b []float64) (float64, error) {
    if len(a) != len(b) {
        return 0, VectorsOfDifferentLen
    }

    res := 0.0

    for i := range len(a) {
        res += a[i] * b[i]
    }

    return res, nil
}

type vectorsOfDifferentLen struct {}
func (vectorsOfDifferentLen ) Error() string {
    return "Vectors must be of the same length";
}
var VectorsOfDifferentLen = vectorsOfDifferentLen{}
