package authorisation

// Checks whether a list contains a particular integer or not
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// String Convert string to *string
func String(v string) *string {
	return &v
}
