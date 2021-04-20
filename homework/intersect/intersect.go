package intersect

// На вход подается два массива произвольной длинны
// Метод должен возвращать пересечение этих массивов (одинаковые элементы в обоих массивах)
// Есть ли способ сделать оптимальнее?
func SliceIntersect(a, b []int64) []int64 {
	resSize := len(a)
	if len(a) > len(b) {
		resSize = len(b)
	}

	result := make([]int64, 0, resSize)
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
