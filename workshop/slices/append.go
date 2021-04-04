package slices

type Search struct {
	userID int64
	query string
	updates bool
}

func copyList(in []string) []string {
	var out []string
	for _, s := range in {
		out = append(out, s)
	}

	return out
}

// Это пример для презентации
func copyListFixed(in []string) []string {
	out := make([]string, len(in))
	for i, s := range in {
		out[i] = s
	}

	return out
}


