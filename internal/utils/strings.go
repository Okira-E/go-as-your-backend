package utils

// @ChatGPT
func toSnakeCase(s string) string {
	var result string
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i != 0 {
				result += "_"
			}
			result += string(r + 32)
		} else {
			result += string(r)
		}
	}
	return result
}
