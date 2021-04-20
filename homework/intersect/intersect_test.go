package intersect

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestSliceIntersect(t *testing.T) {
	tableTests := []struct {
		name string
		sliceA []int64
		sliceB []int64
		result []int64
	} {
		{
			name:   "simple case",
			sliceA: []int64{1, 2, 3, 2, 0},
			sliceB: []int64{5, 1, 2, 7, 3, 2},
			result: []int64{1, 2, 3, 2},
		},
		{
			name:   "not intersect",
			sliceA: []int64{1, 2, 3},
			sliceB: []int64{4, 5, 6},
			result: []int64{},
		},
		{
			name: "second case",
			sliceA: []int64{1, 2, 3},
			sliceB: []int64{1, 2, 3},
			result: []int64{1, 2, 3},
		},
		{
			name: "third case",
			sliceA: []int64{1, 2, 3},
			sliceB: []int64{4, 2, 5, 3, 1},
			result: []int64{2, 3, 1},
		},
	}

	for _, tc := range tableTests {
		t.Run(tc.name, func(t *testing.T) {
			result := BestSliceIntersect(tc.sliceA, tc.sliceB)
			require.Equal(t, tc.result, result)
		})
	}
}

//func BenchmarkSliceIntersect(b *testing.B) {
//	a := getSlice(100)
//	c := getSlice(100)
//	for i := 0; i < b.N; i++ {
//		SliceIntersect(a, c)
//	}
//}

func BenchmarkSlice(b *testing.B) {
	a := getSlice(10000)
	c := getSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BestSliceIntersect(a, c)
	}
}

func getSlice(n int) []int64 {
	res := make([]int64, n)
	for i := 0; i < n; i++ {
		res[i] = rand.Int63n(int64(n))
	}
	return res
}
