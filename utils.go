package kcs

func charMultiplier(n int, char string) string {
	var str = ""
	for i := 0; i< n; i++ {
		str += char
	}
	return str
}
