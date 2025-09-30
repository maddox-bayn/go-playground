package main

// Reverse returns the reverse of the given string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string reads the same forward and backward.
func IsPalindrome(s string) bool {
	runes := []rune(s)
	i, j := 0, len(runes)-1
	for i < j {
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// CharCount counts how many times each rune appears in a string.
func CharCount(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	return counts
}

// RuneCount return the number of runes in a string.
func RuneCount(s string) int {
	runes := []rune(s)
	return len(runes)
}

// FirstRune return the first rune in a string.
func FirstRune(s string) rune {
	return []rune(s)[0]
}

// LastRune return the last rune in a string.
func LastRune(s string) rune {
	return []rune(s)[len([]rune(s))-1]
}
