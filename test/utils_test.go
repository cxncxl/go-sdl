package test

import (
	"testing"
	"time"

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

func TestMeasureTime(t *testing.T) {
    targetDur := time.Millisecond * 200

    _, dur := utils.MeasureTime(func() error {
        timeout := time.NewTicker(targetDur)

        <- timeout.C

        return nil
    })

    // max 1% wrong
    if dur.Microseconds() - targetDur.Microseconds() > targetDur.Microseconds() / 100 {
        t.Errorf(
            "MeasureTime is wrong by %f%%",
            (float64(dur.Microseconds() - targetDur.Microseconds()) / float64(targetDur.Microseconds()) * 100.0),
        )
    }
}
