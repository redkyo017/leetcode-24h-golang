package main

// 3. Longest Substring Without Repeating Characters
func LengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	l := 0
	res := 0
	window := make(map[byte]int)
	for r := 0; r < len(s); r++ {
		if index, ok := window[s[r]]; ok && index >= l {
			l = index + 1
		}
		res = max(res, r-l+1)
		window[s[r]] = r
	}
	return res
}
