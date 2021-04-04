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
