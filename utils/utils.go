package utils

func ContainsString(slice []string, search string) bool {
	for _, item := range slice {
		if item == search {
			return true
		}
	}

	return false
}
