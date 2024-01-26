package main

func linearSearch(haystack []int32, needle int32) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}
