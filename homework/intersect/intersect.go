package intersect

func SliceIntersect(a, b []int64) []int64 {
	cache := make(map[int64]int64)
	for _, val := range a {
		if _, ok := cache[val]; ok {
			cache[val]++
		} else {
			cache[val] = 1
		}
	}

	var result []int64

	for _, valB := range b {
		if num, ok := cache[valB]; ok && num > 0 {
			result = append(result, valB)
			cache[valB]--
		}
	}

	return result
}

func BestSliceIntersect(a, b []int64) []int64 {
	cache := make(map[int64]int64)
	for _, val := range a {
		if _, ok := cache[val]; ok {
			cache[val]++
		} else {
			cache[val] = 1
		}
	}

	result := make([]int64, len(a) + len(b))

	i := 0
	for _, valB := range b {
		if num, ok := cache[valB]; ok && num > 0 {
			i++
			result[i] = valB
			cache[valB]--
		}
	}

	return result[1:i+1]
}

func SliceIntersectSimple(a, b []int64) []int64 {
	var result []int64
	out:
	for _, valA := range a {
		for _, valB := range b {
			if valA == valB {
				result = append(result, valA)
				continue out
			}
		}
	}
	return result
}
