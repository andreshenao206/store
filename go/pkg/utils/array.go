package utils

// Contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, e := range s {
		if e == str {
			return true
		}
	}
	return false
}