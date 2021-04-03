package slices

type Search struct {
	userID int64
	query string
	updates bool
}

func

func copyList(in []string) []string {
	var out []string
	for _, s := range in {
		out = append(out, s)
	}

	return out
}
