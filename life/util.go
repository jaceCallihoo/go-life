package life

import (
   "cmp"
)

func copy2d[T any](dest, src [][]T) int {
    var numCoppied = 0

    for i := range src {
        dest[i] = make([]T, len(src[i]))
        numCoppied += copy(dest[i], src[i])
    }

    return numCoppied
}

func Ptr[T any](val T) *T {
    return &val
}

func Fracture[T any](src []T, pieces int) [][]T {
    var fractured = make([][]T, pieces)
    var pieceLength = len(src) / pieces

    for i := range fractured {
        var pieceStart = i * pieceLength
        var pieceEnd = pieceStart + pieceLength
        fractured[i] = src[pieceStart:pieceEnd]
    }

    return fractured
}

func Reflected[T any](matrix [][]T) [][]T {
    var ret = make([][]T, len(matrix))
    for i := range ret {
        ret[i] = make([]T, len(matrix[i]))
    }

    for i := range matrix {
        for j := 0; j < len(ret[i]) / 2; j++ {
            ret[i][j], ret[i][len(matrix[i])-j-1] = matrix[i][len(matrix[i])-j-1], matrix[i][j]
        }
    }

    return ret
}

func equalSlice[T comparable](a, b []T) bool {
    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}

func equalSlice2d[T comparable](a, b [][]T) bool {
    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if !equalSlice(a[i], b[i]) {
            return false
        }
    }

    return true
}

func Max[T cmp.Ordered](a, b T) T {
    if a > b {
        return a
    }

    return b
}

func Min[T cmp.Ordered](a, b T) T {
    if a < b {
        return a
    }

    return b
}

