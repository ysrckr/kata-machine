package main

import "math"

func binarySearch(haystack *[]int32, needle int32) bool {
	low := 0
	high := len(*haystack)

	for low < high {
		mid := int(math.Floor(float64(low + (high-low)/2)))
		value := (*haystack)[mid]
    
    if (value == needle) {
      return true
    }

		if needle < value {
			high = mid
		} else {
			low = mid + 1
		}

	}

	return false
}
