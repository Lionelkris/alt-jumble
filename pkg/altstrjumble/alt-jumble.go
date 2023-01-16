package altstrjumble

func AltJumble(first, second string) string {
	var res string
	small, large := getSmallAndLargeString(first, second)
	var i int
	for i = 0; i < len(small); i++ {
		res = res + string(small[i]) + string(large[i])
	}
	return res + string(large[i:])
}

func getSmallAndLargeString(f, s string) (string, string) {
	if len(f) > len(s) {
		return s, f
	}
	return f, s
}
