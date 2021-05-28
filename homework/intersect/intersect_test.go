package intersect

import (
	"math/rand"
	"testing"
)

//func TestSliceIntersect(t *testing.T) {
//	tableTests := []struct {
//		name string
//		sliceA []int64
//		sliceB []int64
//		result []int64
//	} {
//		{
//			name:   "simple case",
//			sliceA: []int64{1, 2, 3, 2, 0},
//			sliceB: []int64{5, 1, 2, 7, 3, 2},
//			result: []int64{1, 2, 3, 2},
//		},
//		{
//			name:   "not intersect",
//			sliceA: []int64{1, 2, 3},
//			sliceB: []int64{4, 5, 6},
//			result: nil,
//		},
//		{
//			name: "second case",
//			sliceA: []int64{1, 2, 3},
//			sliceB: []int64{1, 2, 3},
//			result: []int64{1, 2, 3},
//		},
//		{
//			name: "third case",
//			sliceA: []int64{1, 2, 3},
//			sliceB: []int64{4, 2, 5, 3, 1},
//			result: []int64{1, 2, 3},
//		},
//	}
//
//	for _, tc := range tableTests {
//		t.Run(tc.name, func(t *testing.T) {
//			result := SliceIntersectMap(tc.sliceA, tc.sliceB)
//			require.Equal(t, tc.result, result)
//		})
//	}
//}

func BenchmarkIntersect(b *testing.B) {
	a := fillArrays(10000)
	c := fillArrays(10000)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceIntersectMap(a, c)
	}
}

func fillArrays(n int) []int64 {
	result := make([]int64, n)
	for i := 0; i < n; i++ {
		result[i] = rand.Int63()
	}
	return result
}
