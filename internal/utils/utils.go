package utils

import (
	"fmt"
	"time"

	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Ordered] (arr[] T) []T {
    cpy := make([]T, len(arr))
    copy(cpy, arr)

    qs(cpy, 0, len(arr) - 1)

    return cpy
}

func qs[T constraints.Ordered] (arr []T, first int, last int) {
    if first >= last {
        return
    }

    left, right := first, last
    mid := arr[(left + right) / 2]

    for left <= right {
        for arr[left] < mid {
            left += 1
        }
        for arr[right] > mid {
            right -= 1
        }

        if left <= right {
            arr[left], arr[right] = arr[right], arr[left]

            left += 1
            right -= 1
        }
    }

    qs(arr, first, right)
    qs(arr, left, last)
}

func Map[T any, K any] (arr []T, predicate func(v T, i int) K) []K {
    res := make([]K, len(arr))

    for i, v := range arr {
        res[i] = predicate(v, i)
    }

    return res
}

func ForEach[T any] (arr []T, predicate func(v T, i int)) {
    for i, v := range arr {
        predicate(v, i)
    }
}

func Filter[T any] (arr []T, predicate func(v T, i int) bool) []T {
    res := []T{}

    for i, v := range arr {
        if predicate(v, i) {
            res = append(res, v)
        }
    }

    return res
}

func IndexOf[T constraints.Ordered] (arr []T, el T) int {
    for i, v := range arr {
        if el == v {
            return i
        }
    }

    return -1
}

func MeasureTime[T any] (target func() T) (T, time.Duration) {
    t := time.Now()

    res := target()

    dt := time.Now().Sub(t)

    fmt.Println("Time consumed:", dt)

    return res, time.Millisecond / dt
}
