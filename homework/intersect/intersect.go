package intersect

// На вход подается два массива произвольной длинны
// Метод должен возвращать пересечение этих массивов (одинаковые элементы в обоих массивах)
// Есть ли способ сделать оптимальнее?
func SliceIntersect(a, b []int64) []int64 {
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

func SliceIntersectMap(a, b []int64) []int64 {
	m := make(map[int64]struct{}, len(a))
	for _, item := range a {
		m[item] = struct{}{}
	}

	result := make([]int64, 0, len(a))
	for _, item := range b {
		if _, ok := m[item]; ok {
			result = append(result, item)
		}
	}

	return result
}
