package utils

func NonEmptyString(v ...string) string {
	for _, item := range v {
		if item != "" {
			return item
		}
	}
	return ""
}
