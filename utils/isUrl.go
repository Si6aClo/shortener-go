package utils

func IsUrl(url string) bool {
	if url[:7] == "http://" || url[:8] == "https://" {
		return true
	}
	return false
}
