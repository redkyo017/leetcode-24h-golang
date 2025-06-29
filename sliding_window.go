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

// 567. Permutation in String
func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Count := make([]int, 26)
	s2Count := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}
	matches := 0
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}
	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}
		index := s2[r] - 'a'
		s2Count[index]++
		if s2Count[index] == s1Count[index] {
			matches++
		} else if s2Count[index] == s1Count[index]+1 {
			matches--
		}

		leftIndex := s2[l] - 'a'
		s2Count[leftIndex]--
		if s2Count[leftIndex] == s1Count[leftIndex] {
			matches++
		} else if s2Count[leftIndex] == s1Count[leftIndex]-1 {
			matches--
		}
		l++
	}
	return matches == 26
}

// 424. Longest Repeating Character Replacement
func CharacterReplacement(s string, k int) int {
	l := 0
	countChar := make(map[byte]int)
	maxFrequentChar := 0
	res := 0
	for r := range s {
		countChar[s[r]]++
		if countChar[s[r]] > maxFrequentChar {
			maxFrequentChar = countChar[s[r]]
		}

		if (r-l+1)-maxFrequentChar > k {
			countChar[s[l]]--
			l++
		}

		res = max(res, r-l+1)
	}
	return res
}
