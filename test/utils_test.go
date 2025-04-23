package test

import (
	"testing"

	"github.com/cxncxl/gogame/internal/utils"
)

func TestQuickSort(t *testing.T) {
    arr := []int{ 4, 5, 2, 8, 1, 0, 3, 6, 9, 7 }

    sorted := utils.QuickSort(arr)

    for i := range arr {
        if (sorted[i] != i) {
            t.Errorf("Sort didn't work! Expected: %d, got: %d", i, sorted[i])
        }
    }
}
