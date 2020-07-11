package gofaker

// Exclude produces a laundered list of the excluded values
func Exclude(list []string, exclusion map[string]bool) []string {
	var launderedList []string
	for _, v := range list {
		if _, ok := exclusion[v]; !ok {
			launderedList = append(launderedList, v)
		}
	}
	// it the blacklist is a superset of list
	if launderedList == nil {
		return []string{""}
	}
	return launderedList
}
